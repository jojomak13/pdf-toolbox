package core

import (
	"os"
	"path"
)

func (tb *ToolBox) HTML(htmlContent string) error {
	tb.newDirectory()

	chrome := NewChromeInstance()
	chrome.LoadConfig(chrome.GetDefaultConfig())

	defer chrome.CloseContext()

	buffer, err := chrome.Print(htmlContent)
	if err != nil {
		return err
	}

	return os.WriteFile(path.Join(tb.dirPath, os.Getenv("OUTPUT_FILE_NAME")), buffer, 0644)
}
