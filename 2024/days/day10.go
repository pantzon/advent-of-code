package days

import (
	h "aoc/helpers"
	"fmt"
	"path/filepath"
)

type D10Map [][]rune

func reducer(acc D10Map, line string) D10Map {
	var arr []rune
	for _, c := range line {
		arr = append(arr, c)
	}
	return append(acc, arr)
}

func p(x, y int) h.Point {
	return h.Point{X: x, Y: y}
}

type TrailDesignation struct {
	Rating int
	Ends   map[h.Point]bool
}

func trailsFromHead(x int, y int, m D10Map, mem map[h.Point]TrailDesignation) TrailDesignation {
	if res, ok := mem[p(x, y)]; ok {
		return res
	}
	result := TrailDesignation{
		Rating: 0,
		Ends:   map[h.Point]bool{},
	}
	if m[y][x] == '9' {
		result.Rating = 1
		result.Ends[p(x, y)] = true
	} else {
		diffs := [4]h.Point{p(-1, 0), p(1, 0), p(0, -1), p(0, 1)}
		for _, dP := range diffs {
			newY := y + dP.Y
			if 0 <= newY && newY < len(m) {
				newX := x + dP.X
				if 0 <= newX && newX < len(m[newY]) {
					if m[newY][newX]-m[y][x] == 1 {
						subTrails := trailsFromHead(newX, newY, m, mem)
						result.Rating += subTrails.Rating
						for k, v := range subTrails.Ends {
							result.Ends[k] = v
						}
					}
				}
			}
		}
	}
	mem[p(x, y)] = result
	return result
}

func d10Part1(path string) {
	m, err := h.ReduceFile(h.ReduceFileOptions[D10Map]{
		Path:    path,
		Reducer: reducer,
	})
	if err != nil {
		panic(err)
	}

	total := 0
	for y, r := range m {
		for x, c := range r {
			if c == '0' {
				total += len(trailsFromHead(x, y, m, map[h.Point]TrailDesignation{}).Ends)
			}
		}
	}
	fmt.Printf("Sum of Trails: %d\n", total)
}

func d10Part2(path string) {
	m, err := h.ReduceFile(h.ReduceFileOptions[D10Map]{
		Path:    path,
		Reducer: reducer,
	})
	if err != nil {
		panic(err)
	}

	total := 0
	for y, r := range m {
		for x, c := range r {
			if c == '0' {
				total += trailsFromHead(x, y, m, map[h.Point]TrailDesignation{}).Rating
			}
		}
	}
	fmt.Printf("Sum of Ratings: %d\n", total)
}

func Day10() {
	inputFile, err := filepath.Abs("./inputs/day10.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 10")
	d10Part1(inputFile)
	d10Part2(inputFile)
}
