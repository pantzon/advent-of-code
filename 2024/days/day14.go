package days

import (
	h "aoc/helpers"
	"fmt"
	"path/filepath"
	"strings"
)

type BathroomRobot struct {
	Pos      h.Point
	Quadrant int
	V        h.Point
}

func (r *BathroomRobot) Move(ticks int, bounds h.Point) {
	r.Pos = h.Point{X: (r.Pos.X + ticks*r.V.X) % bounds.X, Y: (r.Pos.Y + ticks*r.V.Y) % bounds.Y}
	if r.Pos.X < 0 {
		r.Pos.X += bounds.X
	}
	if r.Pos.Y < 0 {
		r.Pos.Y += bounds.Y
	}
	xCoord := r.Pos.X - bounds.X/2
	yCoord := r.Pos.Y - bounds.Y/2
	r.Quadrant = -1
	if xCoord > 0 {
		if yCoord > 0 {
			r.Quadrant = 0
		} else if yCoord < 0 {
			r.Quadrant = 1
		}
	} else if xCoord < 0 {
		if yCoord < 0 {
			r.Quadrant = 2
		} else if yCoord > 0 {
			r.Quadrant = 3
		}
	}
}

type D14Map struct {
	Bounds h.Point
	Robots []*BathroomRobot
}

func d14Data(path string) D14Map {
	data, err := h.ReduceFile(h.ReduceFileOptions[D14Map]{
		Path: path,
		Reducer: func(acc D14Map, line string) D14Map {
			if len(line) == 0 {
				return acc
			}
			pieces := strings.Split(line, " ")
			robot := BathroomRobot{}
			for _, piece := range pieces {
				r := piece[0]
				coords := strings.Split(piece[2:], ",")
				p := h.Point{X: h.ParseInt(coords[0]), Y: h.ParseInt(coords[1])}
				if r == 'd' {
					acc.Bounds = p
				} else if r == 'p' {
					robot.Pos = p
				} else if r == 'v' {
					robot.V = p
					acc.Robots = append(acc.Robots, &robot)
				}
			}
			return acc
		},
		InitialValue: D14Map{},
	})
	if err != nil {
		panic(err)
	}
	return data
}

func Move(ticks int, data D14Map) {
	m := map[h.Point]int{}
	for _, r := range data.Robots {
		r.Move(ticks, data.Bounds)
		m[r.Pos] += 1
	}
	y := 0
	for y < data.Bounds.Y {
		x := 0
		for x < data.Bounds.X {
			i := m[h.Point{X: x, Y: y}]
			if i == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%d", i)
			}
			x++
		}
		fmt.Printf("\n")
		y++
	}
}

func Quads(data D14Map) [4]int {
	quads := [4]int{}
	for _, r := range data.Robots {
		if r.Quadrant >= 0 {
			quads[r.Quadrant]++
		}
	}
	return quads
}

func SafetyFactor(data D14Map) int {
	quads := Quads(data)
	return quads[0] * quads[1] * quads[2] * quads[3]
}

func d14Part1(path string) {
	m := d14Data(path)
	Move(100, m)
	fmt.Printf("Safety Factor 100: %d\n\n", SafetyFactor(m))
}

func d14Part2(path string) {
	m := d14Data(path)
	// Visually checked, found that they group horizontally and vertically occasionally:
	// Vertically grouped first at 98, then every 103 moves
	// Horizontally grouped firt at 22, then every 101 moves
	// The first time they're grouped in both ways is at:
	//   6587
	Move(6587, m)
}

func Day14() {
	inputFile, err := filepath.Abs("./inputs/day14.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 14")
	d14Part1(inputFile)
	d14Part2(inputFile)
}
