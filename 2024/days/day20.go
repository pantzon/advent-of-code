package days

import (
	h "aoc/helpers"
	"fmt"
	"path/filepath"
	"slices"
	"strings"
)

type D20Map struct {
	Map   [][]rune
	Start h.Point
	End   h.Point
}

func d20Data(path string) D20Map {
	data, err := h.ReduceFile(h.ReduceFileOptions[D20Map]{
		Path: path,
		Reducer: func(acc D20Map, line string) D20Map {
			y := len(acc.Map)
			if strings.ContainsRune(line, 'S') {
				acc.Start = h.Point{X: strings.IndexRune(line, 'S'), Y: y}
			}
			if strings.ContainsRune(line, 'E') {
				acc.End = h.Point{X: strings.IndexRune(line, 'E'), Y: y}
			}
			acc.Map = append(acc.Map, []rune(line))
			return acc
		},
		InitialValue: D20Map{},
	})
	if err != nil {
		panic(err)
	}
	return data
}

func (d D20Map) findRoute() []h.Point {
	route := []h.Point{d.Start}
	for true {
		last := route[len(route)-1]
		if last == d.End {
			break
		}
		for _, n := range [4]h.Point{last.Up(), last.Right(), last.Down(), last.Left()} {
			if d.Map[n.Y][n.X] != '#' && !slices.Contains(route, n) {
				route = append(route, n)
				break
			}
		}
	}
	return route
}

type Cheat struct {
	S h.Point
	E h.Point
}

func nOfNs(p h.Point) map[h.Point]bool {
	return map[h.Point]bool{
		p.Up().Up():       true,
		p.Up().Right():    true,
		p.Right().Right(): true,
		p.Right().Down():  true,
		p.Down().Down():   true,
		p.Down().Left():   true,
		p.Left().Left():   true,
		p.Left().Up():     true,
	}
}

func findCheats(route []h.Point) map[Cheat]int {
	cheats := map[Cheat]int{}
	for i, s := range route {
		ns := nOfNs(s)
		for j, e := range route[i+1:] {
			if ns[e] {
				cheats[Cheat{S: s, E: e}] = j - 1
			}
		}
	}
	return cheats
}

func d20Part1(path string) {
	m := d20Data(path)
	cheats := findCheats(m.findRoute())
	total := 0
	for _, score := range cheats {
		if score >= 100 {
			total++
		}
	}
	fmt.Println("Better than 100:", total)
}

func nOfNs20(p h.Point) map[h.Point]int {
	ns := map[h.Point]int{p: 0}
	lastNs := []h.Point{p}
	i := 0
	for i < 20 {
		i++
		var newLast []h.Point
		for _, last := range lastNs {
			for _, n := range [4]h.Point{last.Up(), last.Right(), last.Down(), last.Left()} {
				if n.X > 0 && n.Y > 0 {
					if _, ok := ns[n]; !ok {
						ns[n] = i
						newLast = append(newLast, n)
					}
				}
			}
		}
		lastNs = newLast
	}
	return ns
}

func findCheats20(route []h.Point) map[Cheat]int {
	cheats := map[Cheat]int{}
	for i, s := range route {
		ns := nOfNs20(s)
		for j, e := range route[i+1:] {
			if t, ok := ns[e]; ok && t > 1 {
				cheats[Cheat{S: s, E: e}] = j - t + 1
			}
		}
	}
	return cheats
}

func d20Part2(path string) {
	m := d20Data(path)
	cheats := findCheats20(m.findRoute())
	total := 0
	for _, score := range cheats {
		if score >= 100 {
			total++
		}
	}
	fmt.Println("Long Cheats better than 100:", total)
}

func Day20() {
	inputFile, err := filepath.Abs("./inputs/day20.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 20")
	d20Part1(inputFile)
	d20Part2(inputFile)
}
