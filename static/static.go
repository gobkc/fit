package static

import "embed"

var (
	//go:embed templates/*
	Files    embed.FS
	filePath = "templates/"
	//go:embed dist/*
	Web     embed.FS
	webPath = "dist/"
)

func GetFileByte(fileName string) (data []byte, err error) {
	data, err = Files.ReadFile(filePath + fileName)
	return
}

func GetWebByte(fileName string) (data []byte, err error) {
	data, err = Web.ReadFile(webPath + fileName)
	return
}
