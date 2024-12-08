package days

import (
	"aoc/helpers"
	"fmt"
	"path/filepath"
)

type Point struct {
	X, Y int
}

func (p Point) InMap(width, height int) bool {
	return 0 <= p.X && p.X < width && 0 <= p.Y && p.Y < height
}

type D8Map struct {
	Antennas map[rune][]Point
	Width    int
	Height   int
}

func d8Process(path string) D8Map {
	mapData, _ := helpers.ReduceFile(helpers.ReduceFileOptions[D8Map]{
		Path: path,
		Reducer: func(acc D8Map, v string) D8Map {
			for i, c := range v {
				if c != '.' {
					acc.Antennas[c] = append(acc.Antennas[c], Point{X: i, Y: acc.Height})
				}
			}
			acc.Width = len(v)
			acc.Height += 1
			return acc
		},
		InitialValue: D8Map{Antennas: map[rune][]Point{}, Width: 0, Height: 0},
	})
	return mapData
}

func d8Part1(path string) {
	mapData := d8Process(path)
	spots := map[Point]bool{}
	for _, a := range mapData.Antennas {
		for i, p1 := range a {
			for _, p2 := range a[i+1:] {
				xDiff := p2.X - p1.X
				yDiff := p2.Y - p1.Y
				overlap1 := Point{X: p1.X - xDiff, Y: p1.Y - yDiff}
				if overlap1.InMap(mapData.Width, mapData.Height) {
					spots[overlap1] = true
				}
				overlap2 := Point{X: p2.X + xDiff, Y: p2.Y + yDiff}
				if overlap2.InMap(mapData.Width, mapData.Height) {
					spots[overlap2] = true
				}
			}
		}
	}
	fmt.Printf("Spots: %d\n", len(spots))
}

func d8Part2(path string) {
	mapData := d8Process(path)
	spots := map[Point]bool{}
	for _, a := range mapData.Antennas {
		for i, p1 := range a {
			if len(a) != 1 {
				spots[p1] = true
			}
			for _, p2 := range a[i+1:] {
				xDiff := p2.X - p1.X
				yDiff := p2.Y - p1.Y
				overlap := Point{X: p1.X - xDiff, Y: p1.Y - yDiff}
				for overlap.InMap(mapData.Width, mapData.Height) {
					spots[overlap] = true
					overlap = Point{X: overlap.X - xDiff, Y: overlap.Y - yDiff}
				}
				overlap = Point{X: p2.X + xDiff, Y: p2.Y + yDiff}
				for overlap.InMap(mapData.Width, mapData.Height) {
					spots[overlap] = true
					overlap = Point{X: overlap.X + xDiff, Y: overlap.Y + yDiff}
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
