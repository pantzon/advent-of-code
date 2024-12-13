package days

import (
	h "aoc/helpers"
	"fmt"
	"path/filepath"
)

type D8Map struct {
	Antennas map[rune][]h.Point
	Width    int
	Height   int
}

func d8Process(path string) D8Map {
	mapData, _ := h.ReduceFile(h.ReduceFileOptions[D8Map]{
		Path: path,
		Reducer: func(acc D8Map, v string) D8Map {
			for i, c := range v {
				if c != '.' {
					acc.Antennas[c] = append(acc.Antennas[c], h.Point{X: i, Y: acc.Height})
				}
			}
			acc.Width = len(v)
			acc.Height += 1
			return acc
		},
		InitialValue: D8Map{Antennas: map[rune][]h.Point{}, Width: 0, Height: 0},
	})
	return mapData
}

func d8Part1(path string) {
	mapData := d8Process(path)
	spots := map[h.Point]bool{}
	for _, a := range mapData.Antennas {
		for i, p1 := range a {
			for _, p2 := range a[i+1:] {
				xDiff := p2.X - p1.X
				yDiff := p2.Y - p1.Y
				overlap1 := h.Point{X: p1.X - xDiff, Y: p1.Y - yDiff}
				if overlap1.InBounds(mapData.Width, mapData.Height) {
					spots[overlap1] = true
				}
				overlap2 := h.Point{X: p2.X + xDiff, Y: p2.Y + yDiff}
				if overlap2.InBounds(mapData.Width, mapData.Height) {
					spots[overlap2] = true
				}
			}
		}
	}
	fmt.Printf("Spots: %d\n", len(spots))
}

func d8Part2(path string) {
	mapData := d8Process(path)
	spots := map[h.Point]bool{}
	for _, a := range mapData.Antennas {
		for i, p1 := range a {
			if len(a) != 1 {
				spots[p1] = true
			}
			for _, p2 := range a[i+1:] {
				xDiff := p2.X - p1.X
				yDiff := p2.Y - p1.Y
				overlap := h.Point{X: p1.X - xDiff, Y: p1.Y - yDiff}
				for overlap.InBounds(mapData.Width, mapData.Height) {
					spots[overlap] = true
					overlap = h.Point{X: overlap.X - xDiff, Y: overlap.Y - yDiff}
				}
				overlap = h.Point{X: p2.X + xDiff, Y: p2.Y + yDiff}
				for overlap.InBounds(mapData.Width, mapData.Height) {
					spots[overlap] = true
					overlap = h.Point{X: overlap.X + xDiff, Y: overlap.Y + yDiff}
				}
			}
		}
	}
	fmt.Printf("Spots: %d\n", len(spots))
}

func Day8() {
	inputFile, err := filepath.Abs("./inputs/day8.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 8")
	d8Part1(inputFile)
	d8Part2(inputFile)
}
