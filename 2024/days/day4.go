package days

import (
	"fmt"
	"path/filepath"
)

func d4Part1(path string) {}

func d4Part2(path string) {}

func Day4() {
	inputFile, err := filepath.Abs("./inputs/day4.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 4")
	d4Part1(inputFile)
	d4Part2(inputFile)
}
