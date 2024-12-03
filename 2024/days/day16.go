package days

import (
	"fmt"
	"path/filepath"
)

func d16Part1(path string) {}

func d16Part2(path string) {}

func Day16() {
	inputFile, err := filepath.Abs("./inputs/day16.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 16")
	d16Part1(inputFile)
	d16Part2(inputFile)
}
