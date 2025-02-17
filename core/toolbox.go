package core

import (
	"io"
	"os"
	"path"
	"sync"
)

type ToolBox struct {
	urls    []string
	dirPath string
	lock    sync.Mutex
}

type IToolBox interface {
	Upload(uploadPath string) (string, error)
}

func NewToolBox(dirName string) *ToolBox {
	return &ToolBox{
		dirPath: path.Join("storage", dirName),
	}
}

func (tb *ToolBox) newDirectory() error {
	return os.Mkdir(tb.dirPath, 0755)
}

func (tb *ToolBox) getFiles() ([]string, error) {
	var res []string

	files, err := os.ReadDir(tb.dirPath)
	if err != nil {
		return res, err
	}

	for _, file := range files {
		res = append(res, path.Join(tb.dirPath, file.Name()))
	}

	return res, nil
}

func (tb *ToolBox) Upload(uploadPath string) (string, error) {
	file, err := os.Open(path.Join(tb.dirPath, os.Getenv("OUTPUT_FILE_NAME")))
	if err != nil {
		return "", err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return UploaderClient.Upload(data, uploadPath)
}

func (tb *ToolBox) Clean() error {
	return os.RemoveAll(tb.dirPath)
}
