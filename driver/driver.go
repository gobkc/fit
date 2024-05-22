package driver

import (
	"github.com/gobkc/fit/conf"
	"time"
)

type Driver struct {
	c *conf.Conf
	NoteDriver
	CompressionDriver
	EmailDriver
}

type NoteDriver interface {
	NewNote(cate, title, content string, upTime ...time.Time) error
	ListCate() (list []string, err error)
	NewCate(cate string) error
	ListNotes(cate string) (list []NoteInstance, err error)
}

type NoteInstance struct {
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	UpdatedTime time.Time `json:"updated_time"`
}

type CompressionDriver interface {
	AddFiles(password string, files ...string) []byte
	DeCompress(password string, file []byte) []CompressionFile
	GetAllFiles(dir string) []string
}

type CompressionFile struct {
	Cate        string    `json:"cate"`
	Filename    string    `json:"filename"`
	Content     string    `json:"content"`
	UpdatedTime time.Time `json:"updated_time"`
}

type EmailDriver interface {
	SendEmail(to string, subject, content string, attachment []byte) error
	GetAttachmentFromEmail() (data []byte, err error)
}
