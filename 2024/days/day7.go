package days

import (
	"fmt"
	"path/filepath"
)

func d7Part1(path string) {}

func d7Part2(path string) {}

func Day7() {
	inputFile, err := filepath.Abs("./inputs/day7.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 7")
	d7Part1(inputFile)
	d7Part2(inputFile)
}
