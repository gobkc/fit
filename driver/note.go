package driver

import (
	"fmt"
	"github.com/gobkc/fit/conf"
	"os"
	"sort"
	"time"
)

type Note struct {
	d Driver
}

func (n *Note) NewNote(cate, title, content string) error {
	conf.IsNotExistCreateCateDir(cate)
	pathSeparator := string(os.PathSeparator)
	notePath := fmt.Sprintf("%s%s%s%s.md", conf.GetCachePath(), cate, pathSeparator, title)
	return os.WriteFile(notePath, []byte(content), 0777)
}

func (n *Note) ListNotes(cate string) (list []NoteInstance, err error) {
	pathSeparator := string(os.PathSeparator)
	dir := conf.GetCachePath() + cate
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %v", err)
	}
	cachePath := dir + pathSeparator
	for _, file := range files {
		if !file.IsDir() {
			content, _ := os.ReadFile(cachePath + file.Name())
			fileInfo, err := os.Stat(cachePath + file.Name())
			var updatedTime time.Time
			if err == nil {
				updatedTime = fileInfo.ModTime()
			}
			list = append(list, NoteInstance{
				Title:       file.Name(),
				Content:     string(content),
				UpdatedTime: updatedTime,
			})
		}
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].UpdatedTime.After(list[j].UpdatedTime)
	})
	return
}

func (n *Note) ListCate() (list []string, err error) {
	var directories []string
	dir := conf.GetCachePath()
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %v", err)
	}

	for _, file := range files {
		if file.IsDir() {
			directories = append(directories, file.Name())
		}
	}

	return directories, nil
}

func (n *Note) NewCate(cate string) error {
	conf.IsNotExistCreateCateDir(cate)
	return nil
}

func NewNote() NoteDriver {
	return &Note{d: d}
}
