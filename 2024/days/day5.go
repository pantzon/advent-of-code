package days

import (
	"fmt"
	"path/filepath"
)

func d5Part1(path string) {}

func d5Part2(path string) {}

func Day5() {
	inputFile, err := filepath.Abs("./inputs/day5.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 5")
	d5Part1(inputFile)
	d5Part2(inputFile)
}
