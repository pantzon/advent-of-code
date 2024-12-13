package days

import (
	h "aoc/helpers"
	"fmt"
	"path/filepath"
	"strings"
)

func d12Data(path string) [][]string {
	data, err := h.ReduceFile(h.ReduceFileOptions[[][]string]{
		Path: path,
		Reducer: func(acc [][]string, line string) [][]string {
			if line == "" {
				return acc
			}
			return append(acc, strings.Split(line, ""))
		}})
	if err != nil {
		panic(err)
	}
	return data
}

type Fence struct {
	// Fence is always from top-left corner of h.Point, going either down or to the right.
	P          h.Point
	IsVertical bool
	IsInside   bool
}

func (f Fence) clone(p h.Point) Fence {
	return Fence{P: p, IsVertical: f.IsVertical, IsInside: f.IsInside}
}

type GardenPlot struct {
	Crop      string
	Size      int
	Perimeter map[Fence]bool
}

func fillPlot(p h.Point, plot *GardenPlot, garden [][]string, visited map[h.Point]*GardenPlot) {
	visited[p] = plot
	r := garden[p.Y]
	plot.Size += 1
	if p.X == 0 || r[p.X-1] != plot.Crop {
		plot.Perimeter[Fence{P: p, IsVertical: true, IsInside: true}] = true
	} else if _, ok := visited[p.Left()]; !ok {
		fillPlot(p.Left(), plot, garden, visited)
	}
	if p.X == len(r)-1 || r[p.X+1] != plot.Crop {
		plot.Perimeter[Fence{P: p.Right(), IsVertical: true, IsInside: false}] = true
	} else if _, ok := visited[p.Right()]; !ok {
		fillPlot(p.Right(), plot, garden, visited)
	}
	if p.Y == 0 || garden[p.Y-1][p.X] != plot.Crop {
		plot.Perimeter[Fence{P: p, IsVertical: false, IsInside: true}] = true
	} else if _, ok := visited[p.Up()]; !ok {
		fillPlot(p.Up(), plot, garden, visited)
	}
	if p.Y == len(garden)-1 || garden[p.Y+1][p.X] != plot.Crop {
		plot.Perimeter[Fence{P: p.Down(), IsVertical: false, IsInside: false}] = true
	} else if _, ok := visited[p.Down()]; !ok {
		fillPlot(p.Down(), plot, garden, visited)
	}
}

func ToGardenPlots(garden [][]string) map[string][]*GardenPlot {
	plots := map[string][]*GardenPlot{}
	visited := map[h.Point]*GardenPlot{}
	for y, r := range garden {
		for x, p := range r {
			point := h.Point{X: x, Y: y}
			if _, ok := visited[point]; ok {
				continue
			}
			plot := &GardenPlot{Crop: p, Size: 0, Perimeter: map[Fence]bool{}}
			fillPlot(point, plot, garden, visited)
			plots[p] = append(plots[p], plot)
			visited[point] = plot
		}
	}
	return plots
}

func d12Part1(path string) {
	data := d12Data(path)
	plots := ToGardenPlots(data)
	total := 0
	for _, pList := range plots {
		for _, plot := range pList {
			total += plot.Size * len(plot.Perimeter)
		}
	}
	fmt.Printf("Total Price: %d\n", total)
}

func calculateSides(plot *GardenPlot) int {
	visited := map[Fence]bool{}
	sides := 0
	for f := range plot.Perimeter {
		if visited[f] {
			continue
		}
		visited[f] = true
		sides += 1
		if f.IsVertical {
			newF := f.clone(f.P.Up())
			for !visited[newF] && plot.Perimeter[newF] {
				visited[newF] = true
				newF = f.clone(newF.P.Up())
			}
			newF = f.clone(f.P.Down())
			for !visited[newF] && plot.Perimeter[newF] {
				visited[newF] = true
				newF = f.clone(newF.P.Down())
			}
		} else {
			newF := f.clone(f.P.Left())
			for !visited[newF] && plot.Perimeter[newF] {
				visited[newF] = true
				newF = f.clone(newF.P.Left())
			}
			newF = f.clone(f.P.Right())
			for !visited[newF] && plot.Perimeter[newF] {
				visited[newF] = true
				newF = f.clone(newF.P.Right())
			}
		}
	}
	return sides
}

func d12Part2(path string) {
	data := d12Data(path)
	plots := ToGardenPlots(data)
	total := 0
	for _, pList := range plots {
		for _, plot := range pList {
			total += plot.Size * calculateSides(plot)
		}
	}
	fmt.Printf("Total Price: %d\n", total)
}

func Day12() {
	inputFile, err := filepath.Abs("./inputs/day12.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 12")
	d12Part1(inputFile)
	d12Part2(inputFile)
}
