package logg

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"
)

type Logger struct {
	file *os.File
}

func (l *Logger) ConfigureLogger(path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	var err error
	l.file, err = os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	multi := io.MultiWriter(os.Stdout, l.file)
	handler := slog.NewTextHandler(multi, nil)
	slog.SetDefault(slog.New(handler))
	return nil
}

func (l *Logger) CloseLogger() {
	l.file.Close()
}
