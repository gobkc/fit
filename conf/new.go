package conf

import (
	"fmt"
	gext "github.com/gobkc/ext"
	"log/slog"
	"os"
	"os/user"
	"sync"
)

var once sync.Once
var conf *Conf

func GetConf() *Conf {
	once.Do(func() {
		conf = &Conf{}
		js := gext.Factory[gext.Json]()
		configFile := GetConfigPath()
		if err := js.UnMarshal(configFile, &conf); err != nil {
			slog.Default().Error(`can't load config.json'`, slog.Any(`path`, configFile), slog.Any(`error detail`, err))
		}
	})
	return conf
}

func GetHomeDir() string {
	currentUser, err := user.Current()
	if err != nil {
		slog.Default().Warn(`failed to get user home dir`, slog.String(`error`, err.Error()))
		return ``
	}
	return currentUser.HomeDir
}

func GetConfigPath() string {
	homeDir := GetHomeDir()
	pathSeparator := string(os.PathSeparator)
	isNotExistCreateDir(homeDir + pathSeparator + `.config` + pathSeparator + `.fit`)
	isNotExistCreateDir(homeDir + pathSeparator + `.cache` + pathSeparator + `.fit`)
	return fmt.Sprintf(`%s%s.config%s.fit%sconfig.json`, homeDir, pathSeparator, pathSeparator, pathSeparator)
}

func GetCachePath() string {
	homeDir := GetHomeDir()
	pathSeparator := string(os.PathSeparator)
	isNotExistCreateDir(homeDir + pathSeparator + `.cache` + pathSeparator + `.fit`)
	return fmt.Sprintf(`%s%s.cache%s.fit%s`, homeDir, pathSeparator, pathSeparator, pathSeparator)
}

func IsNotExistCreateCateDir(cate string) {
	homeDir := GetHomeDir()
	pathSeparator := string(os.PathSeparator)
	isNotExistCreateDir(homeDir + pathSeparator + `.cache` + pathSeparator + `.fit` + pathSeparator + cate)
}

func isNotExistCreateDir(dirName string) {
	_, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		os.MkdirAll(dirName, 0777)
	}
}
