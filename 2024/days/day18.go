package days

import (
	h "aoc/helpers"
	"fmt"
	"path/filepath"
	"slices"
	"strings"
)

type D18Map struct {
	Bounds   h.Point
	Start    h.Point
	End      h.Point
	Blocks   map[h.Point]bool
	CurrTime int
}

type D18Input struct {
	Blocks     []h.Point
	TimeToWait int
}

type D18Data struct {
	Map   *D18Map
	Input *D18Input
}

func d18Data(path string) D18Data {
	data, err := h.ReduceFile(h.ReduceFileOptions[D18Data]{
		Path: path,
		Reducer: func(acc D18Data, line string) D18Data {
			pieces := strings.Split(line, ",")
			if len(pieces) == 3 {
				acc.Map.Bounds = h.Point{X: h.ParseInt(pieces[0]), Y: h.ParseInt(pieces[1])}
				acc.Map.Start = h.Point{X: 0, Y: 0}
				acc.Map.End = h.Point{X: acc.Map.Bounds.X - 1, Y: acc.Map.Bounds.Y - 1}
				acc.Input.TimeToWait = h.ParseInt(pieces[2])
			} else if len(pieces) == 2 {
				acc.Input.Blocks = append(acc.Input.Blocks, h.Point{X: h.ParseInt(pieces[0]), Y: h.ParseInt(pieces[1])})
			}
			return acc
		},
		InitialValue: D18Data{
			Map:   &D18Map{Blocks: map[h.Point]bool{}},
			Input: &D18Input{},
		},
	})
	if err != nil {
		panic(err)
	}
	return data
}

func (d D18Data) ProcessTime(ticks int) bool {
	i := 0
	for i < ticks {
		if d.Map.CurrTime > len(d.Input.Blocks) {
			return false
		}
		d.Map.Blocks[d.Input.Blocks[d.Map.CurrTime]] = true
		i++
		d.Map.CurrTime++
	}
	return true
}

func (d D18Data) Print() {
	y := 0
	for y < d.Map.Bounds.Y {
		x := 0
		for x < d.Map.Bounds.X {
			if d.Map.Blocks[h.Point{X: x, Y: y}] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
			x++
		}
		fmt.Println("")
		y++
	}
	fmt.Println("")
}

func (d D18Data) FindPath(locs map[h.Point]bool, visited []h.Point) (int, bool) {
	if locs[d.Map.End] {
		return 0, true
	}
	next := map[h.Point]bool{}
	for loc, _ := range locs {
		for _, p := range [4]h.Point{loc.Right(), loc.Down(), loc.Left(), loc.Up()} {
			if p.InBounds(d.Map.Bounds.X, d.Map.Bounds.Y) &&
				!d.Map.Blocks[p] &&
				!slices.Contains(visited, p) {
				next[p] = true
			}
		}
		visited = append(visited, loc)
	}
	if len(next) == 0 {
		return 0, false
	}
	v, ok := d.FindPath(next, visited)
	return v + 1, ok
}

func d18Part1(path string) {
	data := d18Data(path)
	data.ProcessTime(data.Input.TimeToWait)
	length, _ := data.FindPath(map[h.Point]bool{data.Map.Start: true}, nil)
	fmt.Println("Min Path:", length)
}

func d18Part2(path string) {
	data := d18Data(path)
	data.ProcessTime(data.Input.TimeToWait)
	for {
		_, ok := data.FindPath(map[h.Point]bool{data.Map.Start: true}, nil)
		if !ok {
			fmt.Println("Blocking Block:", data.Input.Blocks[data.Map.CurrTime-1])
			break
		}
		ok = data.ProcessTime(1)
		if !ok {
			fmt.Println("Ran out of Blocks!")
			break
		}
	}
}

func Day18() {
	inputFile, err := filepath.Abs("./inputs/day18.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 18")
	d18Part1(inputFile)
	d18Part2(inputFile)
}
