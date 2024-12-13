package days

import (
	h "aoc/helpers"
	"fmt"
	"math"
	"path/filepath"
	"strings"
)

func d11ParseData(path string) []int {
	data, err := h.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var out []int
	for _, i := range strings.Split(data, " ") {
		out = append(out, h.ParseInt(i))
	}
	return out
}

func getOuts(stone int) []int {
	var outs []int
	if stone == 0 {
		outs = []int{1}
	} else {
		magnitude := int(math.Log10(float64(stone))) + 1
		if magnitude%2 == 0 {
			outs = []int{stone / int(math.Pow10(magnitude/2)), stone % int(math.Pow10(magnitude/2))}
		} else {
			outs = []int{stone * 2024}
		}
	}
	return outs
}

func processStones(stones []int, maxDepth int) map[int][]int {
	uniqueStones := map[int][]int{}
	toProcess := stones
	for maxDepth > 0 {
		var newStones []int
		for _, stone := range toProcess {
			outs := getOuts(stone)
			uniqueStones[stone] = outs
			for _, out := range outs {
				if _, ok := uniqueStones[out]; !ok {
					newStones = append(newStones, out)
				}
			}
		}
		toProcess = newStones
		maxDepth--
	}
	return uniqueStones
}

func recursiveCount(start int, depth int, uniques map[int][]int, memory map[int]map[int]int) int {
	if depth == 0 {
		return 1
	}
	if m, ok := memory[start]; ok {
		if c, ok := m[depth]; ok {
			return c
		}
	} else {
		memory[start] = map[int]int{}
	}
	count := 0
	for _, s := range uniques[start] {
		count += recursiveCount(s, depth-1, uniques, memory)
	}
	memory[start][depth] = count
	return count
}

func countStones(starts []int, depth int) int {
	uniques := processStones(starts, depth)
	memory := map[int]map[int]int{}
	count := 0
	for _, start := range starts {
		count += recursiveCount(start, depth, uniques, memory)
	}
	return count
}

func d11Part1(path string) {
	stones := d11ParseData(path)
	fmt.Printf("Stone count 25: %d\n", countStones(stones, 25))
}

func d11Part2(path string) {
	stones := d11ParseData(path)
	fmt.Printf("Stone count 75: %d\n", countStones(stones, 75))
}

func Day11() {
	inputFile, err := filepath.Abs("./inputs/day11.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 11")
	d11Part1(inputFile)
	d11Part2(inputFile)
}
