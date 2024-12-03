package days

import (
	"fmt"
	"path/filepath"
)

func d22Part1(path string) {}

func d22Part2(path string) {}

func Day22() {
	inputFile, err := filepath.Abs("./inputs/day22.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 22")
	d22Part1(inputFile)
	d22Part2(inputFile)
}
