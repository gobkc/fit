package driver

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/emersion/go-imap/client"
	"github.com/gobkc/to"
	"log"
	"net"
	"net/smtp"
	"os"
	"strings"
	"time"
)

type Email struct {
	d *Driver
}

func NewEmail() EmailDriver {
	return &Email{d: d}
}

func (e *Email) SendEmail(to string, subject, content string, attachment []byte) error {
	conf := e.d.c.Email

	from := conf.User
	password := conf.Pass

	body := content
	hostSplit := strings.Split(conf.Smtp, ":")
	if len(hostSplit) != 2 {
		return fmt.Errorf("SMTP host must be in format <host>:<port>")
	}
	smtpHost, smtpPort := hostSplit[0], hostSplit[1]

	// email headers
	headers := make(map[string]string)
	headers["From"] = from
	headers["To"] = to
	headers["Subject"] = subject

	// build MIME email message
	message := buildMIMEMessage(headers, body, attachment)

	// connect SMTP server
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// send email
	err := smtp.SendMail(fmt.Sprintf("%s:%s", smtpHost, smtpPort), auth, from, []string{to}, []byte(message))
	if err != nil {
		return fmt.Errorf("Failed to send email: %v\n", err)
	}

	fmt.Println("Email sent successfully")
	return nil
}

func (e *Email) GetAttachmentFromEmail() (data []byte, err error) {
	conf := e.d.c.Email
	var server string
	var port int
	if split := strings.Split(conf.Imap, `:`); len(split) != 2 {
		return nil, errors.New("imap host must be in format <host>:<port>")
	} else {
		server = split[0]
		port = to.Int[int](split[1])
	}
	return readEmail(server, port, conf.User, conf.Pass)
}

// buildMIMEMessage build MIME format email content
func buildMIMEMessage(headers map[string]string, body string, attachment []byte) string {
	var message strings.Builder

	// write headers
	for key, value := range headers {
		message.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
	}

	// write MIME email body
	message.WriteString("MIME-Version: 1.0\r\n")
	message.WriteString("Content-Type: multipart/mixed; boundary=boundary2fit\r\n")
	message.WriteString("\r\n")
	message.WriteString("--boundary2fit\r\n")
	message.WriteString("Content-Type: text/plain; charset=utf-8\r\n")
	message.WriteString("\r\n")
	message.WriteString(body)
	message.WriteString("\r\n")

	// add attachment
	message.WriteString("--boundary2fit\r\n")
	message.WriteString("Content-Type: application/octet-stream\r\n")
	message.WriteString("Content-Disposition: attachment; filename=fit.fit\r\n")
	message.WriteString("\r\n")
	message.Write(attachment)
	message.WriteString("\r\n")

	// MIME email end flag
	message.WriteString("--boundary2fit--\r\n")

	return message.String()
}

func getIMAPReaderWriter(server string, port int, username, password string) (*bufio.Reader, *bufio.Writer, net.Conn, error) {
	// Establish a connection to the IMAP server
	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", server, port), nil)
	if err != nil {
		return nil, nil, nil, err
	}

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	// Read server greeting
	_, err = reader.ReadString('\n')
	if err != nil {
		return nil, nil, nil, err
	}

	// Send LOGIN command
	if err := sendCommand(writer, reader, fmt.Sprintf("a001 LOGIN %s %s", username, password)); err != nil {
		return nil, nil, nil, err
	}

	return reader, writer, conn, nil
}

func getAttachment(reader *bufio.Reader, writer *bufio.Writer, emailId int, subject string) (attachment []byte, err error) {
	emailCommand := fmt.Sprintf(`a003 FETCH %d (BODY[HEADER.FIELDS (SUBJECT)] BODY[TEXT])`, emailId)
	if err = sendCommand(writer, reader, emailCommand); err != nil {
		return nil, err
	}

	// Read email content
	var content bytes.Buffer
	fileTag := false
	subjectFlag := false
	var emailSubject string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		if strings.HasPrefix(line, "Subject: ") && subjectFlag == false {
			fmt.Sscanf(line, "Subject: %s", &emailSubject)
			if emailSubject != subject {
				return []byte{}, errors.New(`it's not a valid fit email subject'`)
			}
		}

		if strings.HasPrefix(line, "a003 OK") || strings.HasPrefix(line, "--boundary2fit--") {
			break
		}

		if fileTag {
			content.WriteString(line)
		}

		if strings.HasPrefix(line, "Content-Disposition") {
			fileTag = true
		}
	}

	if emailSubject == "" {
		return nil, errors.New("no fit email subject")
	}
	attachment = content.Bytes()
	if contentLen := len(attachment); contentLen > 4 {
		attachment = attachment[2 : contentLen-2]
	}

	return attachment, nil
}

func readEmail(server string, port int, username, password string) (attachment []byte, err error) {
	reader, writer, conn, err := getIMAPReaderWriter(server, port, username, password)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// Select the INBOX
	if err := sendCommand(writer, reader, "a002 SELECT INBOX"); err != nil {
		return nil, err
	}

	// 发送STATUS命令获取INBOX邮箱的状态信息
	if err := sendCommand(writer, reader, "a002 STATUS INBOX (MESSAGES)"); err != nil {
		return nil, err
	}

	// Send UID SEARCH command to get all message UIDs
	if err := sendCommand(writer, reader, "a003 UID SEARCH ALL"); err != nil {
		return nil, err
	}

	// Read response to UID SEARCH command to get the number of messages
	var numMessages int
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		if strings.HasPrefix(line, "a003 OK") {
			break
		}
		if strings.HasPrefix(line, "* SEARCH") {
			// Count the number of UIDs in the SEARCH response
			parts := strings.Split(line, " ")
			numMessages = len(parts) - 2 // subtract "* SEARCH " prefix
		}
	}

	for i := numMessages; i > 0; i-- {
		attachment, err = getAttachment(reader, writer, i, `fit`)
		if err == nil {
			break
		}
	}

	// Logout
	if err := sendCommand(writer, reader, "a004 LOGOUT"); err != nil {
		return nil, err
	}

	os.WriteFile(`/home/xiong/fit.email`, attachment, 0777)

	return
}

func sendCommand(writer *bufio.Writer, reader *bufio.Reader, command string) error {
	_, err := writer.WriteString(command + "\r\n")
	if err != nil {
		return err
	}
	writer.Flush()
	_, err = reader.ReadString('\n')
	return err
}

//func getAttachmentFromEmail(server, email, password, subject string) ([]byte, error) {
//	c, err := Login(server, email, password)
//	if err != nil {
//		return nil, err
//	}
//
//	idClient := emailClientId.NewClient(c)
//	idClient.ID(
//		emailClientId.ID{
//			emailClientId.FieldName:    "IMAPClient",
//			emailClientId.FieldVersion: "2.1.0",
//		},
//	)
//
//	boxes := make(chan *imap.MailboxInfo, 100)
//	boxesDone := make(chan error, 1)
//	go func() {
//		boxesDone <- c.List("", "*", boxes)
//	}()
//	for box := range boxes {
//		mbox, err := c.Select(box.Name, false)
//		if err != nil {
//			return nil, err
//		}
//		if mbox.Messages == 0 {
//			continue
//		}
//		criteria := imap.NewSearchCriteria()
//		criteria.Since = time.Now().Add(-365 * time.Hour * 24)
//		ids, err := c.UidSearch(criteria)
//		if err != nil {
//			continue
//		}
//		if len(ids) == 0 {
//			continue
//		}
//
//		seqSet := new(imap.SeqSet)
//		seqSet.AddNum(ids...)
//		sect := &imap.BodySectionName{Peek: true}
//		messages := make(chan *imap.Message, 100)
//		messageDone := make(chan error, 1)
//		go func() {
//			messageDone <- c.UidFetch(seqSet, []imap.FetchItem{sect.FetchItem()}, messages)
//		}()
//
//		var latestDate time.Time
//		var latestMsg *imap.Message
//
//		for msg := range messages {
//			r := msg.GetBody(sect)
//			mr, err := mail.CreateReader(r)
//			if err != nil {
//				return nil, err
//			}
//			header := mr.Header
//			title, _ := header.Subject()
//			if title != subject {
//				continue
//			}
//			date, _ := header.Date()
//			if date.After(latestDate) {
//				latestDate = date
//				latestMsg = msg
//			}
//			fmt.Println("读取到:", title, " 在:", date.Format(time.RFC3339))
//		}
//
//		if latestMsg == nil {
//			return nil, fmt.Errorf("no fit email found")
//		}
//
//		r := latestMsg.GetBody(&imap.BodySectionName{Peek: true})
//		if r == nil {
//			return nil, fmt.Errorf("no body found in the email")
//		}
//		buf := new(strings.Builder)
//		_, err = io.Copy(buf, r)
//		if err != nil {
//			return nil, fmt.Errorf("error copying message body: %w", err)
//		}
//		reader := strings.NewReader(buf.String())
//
//		mr, err := mail.CreateReader(reader)
//		if err != nil {
//			return nil, fmt.Errorf("error creating mail reader: %w", err)
//		}
//
//		for {
//			part, err := mr.NextPart()
//			if err == io.EOF {
//				break
//			}
//			if err != nil {
//				return nil, fmt.Errorf("error reading next part: %w", err)
//			}
//
//			switch h := part.Header.(type) {
//			case *mail.AttachmentHeader:
//				filename, _ := h.Filename()
//				if strings.Contains(filename, "fit") { // Replace with your condition
//					attachment, err := io.ReadAll(part.Body)
//					if err != nil {
//						return nil, fmt.Errorf("error reading attachment body: %w", err)
//					}
//					return attachment, nil
//				}
//			}
//		}
//	}
//
//	return nil, nil
//}

func Login(server, user, pass string) (*client.Client, error) {
	dial := new(net.Dialer)
	dial.Timeout = time.Duration(3) * time.Second
	c, err := client.DialWithDialerTLS(dial, server, nil)
	if err != nil {
		c, err = client.DialWithDialer(dial, server)
		if err != nil {
			log.Printf("[ERR]\tLOGIN EMAIL:%s", err.Error())
			return nil, err
		}
	}
	if err = c.Login(user, pass); err != nil {
		log.Printf("[ERR]\tLOGIN EMAIL:%s", err.Error())
	}
	return c, err
}
