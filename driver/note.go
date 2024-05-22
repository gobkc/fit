package driver

import (
	"fmt"
	"github.com/gobkc/fit/conf"
	"os"
	"sort"
	"strings"
	"time"
)

type Note struct {
	d *Driver
}

func (n *Note) NewNote(cate, title, content string, upTime ...time.Time) error {
	conf.IsNotExistCreateCateDir(cate)
	pathSeparator := string(os.PathSeparator)
	notePath := fmt.Sprintf("%s%s%s%s.md", conf.GetCachePath(), cate, pathSeparator, title)
	err := os.WriteFile(notePath, []byte(content), 0777)
	if err != nil {
		return err
	}
	if len(upTime) > 0 {
		return os.Chtimes(notePath, upTime[0], upTime[0])
	}
	return nil
}

func (n *Note) ListNotes(cate string) (list []NoteInstance, err error) {
	pathSeparator := string(os.PathSeparator)
	dir := conf.GetCachePath()
	if cate == `` || cate == `undefined` || cate == `{cate}` {
		files := n.d.GetAllFiles(dir)
		for _, fileName := range files {
			content, _ := os.ReadFile(fileName)
			fileInfo, err := os.Stat(fileName)
			var updatedTime time.Time
			if err == nil {
				updatedTime = fileInfo.ModTime()
			}
			list = append(list, NoteInstance{
				Title:       strings.TrimSuffix(fileInfo.Name(), `.md`),
				Content:     string(content),
				UpdatedTime: updatedTime,
			})
		}
	} else {
		dir += cate
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
					Title:       strings.TrimSuffix(file.Name(), `.md`),
					Content:     string(content),
					UpdatedTime: updatedTime,
				})
			}
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
