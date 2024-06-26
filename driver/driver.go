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
	DeleteCate(cate string) error
	DeleteNote(cate, title string) error
	ListNotes(cate string) (list []NoteInstance, err error)
	ListConfigurations() (list []conf.Conf, mainConf string, err error)
	EnableConfiguration(conf conf.Conf) (err error)
	DeleteConfiguration(fileName string) (err error)
	CreateConfiguration(conf conf.Conf) (err error)
}

type NoteInstance struct {
	Cate        string    `json:"cate"`
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
