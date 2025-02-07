package logger

import (
	"io"
	"log"
	"os"
)

// Init создает и возвращает кастомный логгер
func Init(logFile string) (*log.Logger, error) {
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	// Логирование и в файл, и в консоль
	multiWriter := io.MultiWriter(os.Stdout, file)
	logger := log.New(multiWriter, "", log.LstdFlags|log.Lshortfile)

	return logger, nil
}
