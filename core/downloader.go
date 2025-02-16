package core

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"sync"
)

var mu sync.Mutex

func DownloadFiles(urls []string, dirPath string) error {
	var wg sync.WaitGroup
	errors := make(chan error, len(urls))

	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			if err := DownloadFile(url, dirPath); err != nil {
				errors <- fmt.Errorf("cannot download file [%s]: %v", url, err)
			}
		}(url)
	}

	wg.Wait()
	close(errors)

	// Collect any errors
	var downloadErrors []error
	for err := range errors {
		downloadErrors = append(downloadErrors, err)
		Logger.Printf(err.Error())
	}

	if len(downloadErrors) > 0 {
		return fmt.Errorf("encountered %d errors while downloading images", len(downloadErrors))
	}

	return nil
}

func DownloadFile(url, dirPath string) error {
	filePath := filepath.Join(dirPath, path.Base(url))

	mu.Lock()

	defer mu.Unlock()

	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		return nil
	}

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 response code: %d", response.StatusCode)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)

	return err
}
