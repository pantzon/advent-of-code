package days

import (
	h "aoc/helpers"
	"fmt"
	"path/filepath"
	"regexp"
)

func d3Part1(path string) {
	instructions, err := h.ParseFileToLines(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	re := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)

	total := 0
	for _, line := range instructions {
		for _, match := range re.FindAllStringSubmatch(line, -1) {
			total += h.ParseInt(match[1]) * h.ParseInt(match[2])
		}
	}

	fmt.Printf("Multiple Total: %d\n", total)
}

func d3Part2(path string) {
	instructions, err := h.ParseFileToLines(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	re := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)|do\(\)|don't\(\)`)

	enabled := true
	total := 0
	for _, line := range instructions {
		for _, match := range re.FindAllStringSubmatch(line, -1) {
			if match[0] == "do()" {
				enabled = true
			} else if match[0] == "don't()" {
				enabled = false
			} else if enabled {
				total += h.ParseInt(match[1]) * h.ParseInt(match[2])
			}
		}
	}

	fmt.Printf("Multiple Total: %d\n", total)
}

func Day3() {
	inputFile, err := filepath.Abs("./inputs/day3.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 3")
	d3Part1(inputFile)
	d3Part2(inputFile)
}
