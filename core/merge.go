package core

import (
	"os"
	"path"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

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
