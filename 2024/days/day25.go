package days

import (
	"fmt"
	"path/filepath"
)

func d25Part1(path string) {}

func d25Part2(path string) {}

func Day25() {
	inputFile, err := filepath.Abs("./inputs/day25.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 25")
	d25Part1(inputFile)
	d25Part2(inputFile)
}
