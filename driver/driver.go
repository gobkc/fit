package driver

import (
	"github.com/gobkc/fit/conf"
	"time"
)

type Driver struct {
	c *conf.Conf
	NoteDriver
}

type NoteDriver interface {
	NewNote(cate, title, content string) error
	ListCate() (list []string, err error)
	NewCate(cate string) error
	ListNotes(cate string) (list []NoteInstance, err error)
}

type NoteInstance struct {
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	UpdatedTime time.Time `json:"updated_time"`
}
