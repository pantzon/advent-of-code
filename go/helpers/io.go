package helpers

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

type ParseFileOptions[O any] struct {
	Path   string
	Parser func(line string) O
}

type ReduceFileOptions[O any] struct {
	Path         string
	Reducer      func(acc O, line string) O
	InitialValue O
}

func ReadFile(path string) (string, error) {
	if out, err := ParseFileToLines(path); err == nil {
		return strings.Join(out, "\n"), nil
	} else {
		return "", err
	}
}

func ParseFile[O any](opts ParseFileOptions[O]) ([]O, error) {
	reducer := func(acc []O, line string) []O {
		return append(acc, opts.Parser(line))
	}
	return ReduceFile(ReduceFileOptions[[]O]{
		Path:    opts.Path,
		Reducer: reducer,
	})
}

func ParseFileToLines(path string) ([]string, error) {
	return ParseFile(ParseFileOptions[string]{
		Path:   path,
		Parser: func(line string) string { return line },
	})
}

func ReduceFile[O any](opts ReduceFileOptions[O]) (O, error) {
	acc := opts.InitialValue
	path, err := filepath.Abs(opts.Path)
	if err != nil {
		return acc, err
	}
	file, err := os.Open(path)
	if err != nil {
		return acc, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		acc = opts.Reducer(acc, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return acc, err
	}
	return acc, nil
}
