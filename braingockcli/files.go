package main

import (
	"os"
	"path/filepath"
)

func openFile(path string) (*os.File, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	reader, err := os.Open(abs)
	if err != nil {
		return nil, err
	}
	return reader, nil
}

func createFile(path string) (*os.File, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	writer, err := os.Create(abs)
	if err != nil {
		return nil, err
	}
	return writer, nil
}
