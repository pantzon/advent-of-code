package days

import (
	h "aoc/helpers"
	"fmt"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
)

func d19Data(path string) Towels {
	data, err := h.ReduceFile(h.ReduceFileOptions[Towels]{
		Path: path,
		Reducer: func(acc Towels, line string) Towels {
			pieces := strings.Split(line, ", ")
			if len(pieces) > 1 {
				slices.Sort(pieces)
				acc.Available = pieces
			} else if len(line) > 0 {
				acc.Wanted = append(acc.Wanted, line)
			}
			return acc
		},
		InitialValue: Towels{},
	})
	if err != nil {
		panic(err)
	}
	return data
}

type Towels struct {
	Available []string
	Wanted    []string
}

func (t Towels) PossiblePatterns() []string {
	regex := regexp.MustCompile(fmt.Sprintf("^(%s)+$", strings.Join(t.Available, "|")))
	var possible []string
	for _, pattern := range t.Wanted {
		if regex.Match([]byte(pattern)) {
			possible = append(possible, pattern)
		}
	}
	return possible
}

func d19Part1(path string) {
	towels := d19Data(path)
	possible := towels.PossiblePatterns()
	fmt.Println("Possible Patterns: ", len(possible))
}

func HowManyWays(pattern string, available []string, found map[string]int) int {
	if ways, ok := found[pattern]; ok {
		return ways
	}
	ways := 0
	for _, avail := range available {
		if strings.HasPrefix(pattern, avail) {
			if len(pattern) == len(avail) {
				ways += 1
			} else {
				ways += HowManyWays(pattern[len(avail):], available, found)
			}
		}
		if avail > pattern {
			break
		}
	}
	found[pattern] = ways
	return ways
}

func d19Part2(path string) {
	towels := d19Data(path)
	total := 0
	found := map[string]int{"": 0}
	for _, pattern := range towels.PossiblePatterns() {
		total += HowManyWays(pattern, towels.Available, found)
	}
	fmt.Println("How Many Ways:", total)
}

func Day19() {
	inputFile, err := filepath.Abs("./inputs/day19.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 19")
	d19Part1(inputFile)
	d19Part2(inputFile)
}
