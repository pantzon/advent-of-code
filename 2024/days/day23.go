package days

import (
	"fmt"
	"path/filepath"
)

func d23Part1(path string) {}

func d23Part2(path string) {}

func Day23() {
	inputFile, err := filepath.Abs("./inputs/day23.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 23")
	d23Part1(inputFile)
	d23Part2(inputFile)
}
