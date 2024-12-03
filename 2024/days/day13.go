package days

import (
	"fmt"
	"path/filepath"
)

func d13Part1(path string) {}

func d13Part2(path string) {}

func Day13() {
	inputFile, err := filepath.Abs("./inputs/day13.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 13")
	d13Part1(inputFile)
	d13Part2(inputFile)
}
