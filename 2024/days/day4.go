package days

import (
	h "aoc/helpers"
	"fmt"
	"path/filepath"
)

func checkXmas(data []string, x int, y int, xDiff int, yDiff int) int {
	if 0 <= y+3*yDiff && y+3*yDiff < len(data) &&
		0 <= x+3*xDiff && x+3*xDiff < len(data[y]) &&
		data[y][x] == 'X' &&
		data[y+yDiff][x+xDiff] == 'M' &&
		data[y+2*yDiff][x+2*xDiff] == 'A' &&
		data[y+3*yDiff][x+3*xDiff] == 'S' {
		return 1
	}
	return 0
}

func d4Part1(path string) {
	wordSearch, err := h.ParseFileToLines(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	total := 0
	for i, _ := range wordSearch {
		for j, _ := range wordSearch[i] {
			total += checkXmas(wordSearch, i, j, 0, 1) +
				checkXmas(wordSearch, i, j, 1, 1) +
				checkXmas(wordSearch, i, j, 1, 0) +
				checkXmas(wordSearch, i, j, 1, -1) +
				checkXmas(wordSearch, i, j, 0, -1) +
				checkXmas(wordSearch, i, j, -1, -1) +
				checkXmas(wordSearch, i, j, -1, 0) +
				checkXmas(wordSearch, i, j, -1, 1)
		}
	}
	fmt.Printf("Xmas-es found: %d\n", total)
}

func checkXMas(data []string, x int, y int) int {
	if 0 <= y-1 && y+1 < len(data) &&
		0 <= x-1 && x+1 < len(data[y]) &&
		data[y][x] == 'A' &&
		((data[y+1][x+1] == 'M' && data[y-1][x-1] == 'S') ||
			(data[y+1][x+1] == 'S' && data[y-1][x-1] == 'M')) &&
		((data[y+1][x-1] == 'M' && data[y-1][x+1] == 'S') ||
			(data[y+1][x-1] == 'S' && data[y-1][x+1] == 'M')) {
		return 1
	}
	return 0
}

func d4Part2(path string) {
	wordSearch, err := h.ParseFileToLines(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	total := 0
	for i, _ := range wordSearch {
		for j, _ := range wordSearch[i] {
			total += checkXMas(wordSearch, i, j)
		}
	}
	fmt.Printf("X-Mas-es found: %d\n", total)
}

func Day4() {
	inputFile, err := filepath.Abs("./inputs/day4.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 4")
	d4Part1(inputFile)
	d4Part2(inputFile)
}
