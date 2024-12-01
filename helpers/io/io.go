package io

import (
	"bufio"
	"os"
	"path/filepath"
)

type ReadFileOptions[O any] struct {
	Path   string
	Parser func(line string) O
}

func ReadFile(path string) ([]string, error) {
	return ReadAndParseFile(ReadFileOptions[string]{
		Path:   path,
		Parser: func(line string) string { return line },
	})
}

func ReadAndParseFile[O any](opts ReadFileOptions[O]) ([]O, error) {
	path, err := filepath.Abs(opts.Path)
	if err != nil {
		return nil, err
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []O
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, opts.Parser(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return data, nil
}
