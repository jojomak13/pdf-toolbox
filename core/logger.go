package core

import (
	"fmt"
	"log"
	"os"
)

var Logger *log.Logger

func LoadLogger() {
	Logger = log.New(loadFile(), os.Getenv("APP_NAME")+": ", log.Ldate|log.Ltime)
}

func loadFile() *os.File {
	file, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Sprintf("Failed to open log file: %v", err))
	}

	return file
}
