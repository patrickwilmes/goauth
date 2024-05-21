package common

import (
	"log"
	"os"

	"github.com/rs/zerolog"
)

var file *os.File

func init() {
	file, _ = os.OpenFile(
		"application.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
}

func GetLogger() zerolog.Logger {
	multi := zerolog.MultiLevelWriter(zerolog.ConsoleWriter{Out: os.Stdout}, file)
	logger := zerolog.New(multi).With().Timestamp().Logger()
	return logger
}

func CloseLogging() {
	err := file.Close()
	if err != nil {
		log.Fatalf("Failed to close file handle for logger: %v", err)
	}
}
