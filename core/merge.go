package core

import (
	"os"
	"path"
	"sync"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

type ToolBox struct {
	urls    []string
	dirPath string
	lock    sync.Mutex
}

func NewToolBox(dirName string) *ToolBox {
	return &ToolBox{
		dirPath: path.Join("storage", dirName),
	}
}

func (tb *ToolBox) Merge(urls []string) error {
	tb.newDirectory()

	if err := DownloadFiles(urls, tb.dirPath); err != nil {
		return err
	}

	files, err := tb.getFiles()
	if err != nil {
		return err
	}

	return api.MergeAppendFile(files, path.Join(tb.dirPath, os.Getenv("OUTPUT_FILE_NAME")), false, model.NewDefaultConfiguration())
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
