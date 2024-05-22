package driver

import (
	"github.com/gobkc/fit/conf"
	"sync"
)

var d = &Driver{}
var once sync.Once

func NewDriver() *Driver {
	once.Do(func() {
		d.c = conf.GetConf()
		d.NoteDriver = NewNote()
		d.CompressionDriver = NewCompression()
		d.EmailDriver = NewEmail()
	})
	return d
}
