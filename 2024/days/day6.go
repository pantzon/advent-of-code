package days

import (
	"aoc/helpers"
	"fmt"
	"path/filepath"
	"slices"
	"strings"
)

type D6Data struct {
	Area   [][]rune
	StartX int
	StartY int
}

func getD6Data(path string) *D6Data {
	reducer := func(acc D6Data, line string) D6Data {
		if x := strings.Index(line, "^"); x >= 0 {
			acc.StartX = x
			acc.StartY = len(acc.Area)
		}
		acc.Area = append(acc.Area, []rune(line))
		return acc
	}
	data, err := helpers.ReduceFile(helpers.ReduceFileOptions[D6Data]{
		Path:         path,
		Reducer:      reducer,
		InitialValue: D6Data{},
	})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &data
}

type GuardDirection int

const (
	Up GuardDirection = iota
	Right
	Down
	Left
)

func nextPos(dir GuardDirection, x int, y int) (int, int) {
	switch dir {
	case Up:
		return x, y - 1
	case Right:
		return x + 1, y
	case Down:
		return x, y + 1
	case Left:
		return x - 1, y
	default:
		return -1, -1
	}
}

func turn(dir GuardDirection) GuardDirection {
	return (dir + 1) % 4
}

func guardTravel(data *D6Data) (map[helpers.Point][]GuardDirection, bool) {
	visited := map[helpers.Point][]GuardDirection{}
	guardX := data.StartX
	guardY := data.StartY
	direction := Up
	for inRange := true; inRange; {
		if slices.Contains(visited[helpers.Point{X: guardX, Y: guardY}], direction) {
			return visited, false
		}
		visited[helpers.Point{X: guardX, Y: guardY}] = append(visited[helpers.Point{X: guardX, Y: guardY}], direction)
		for {
			newX, newY := nextPos(direction, guardX, guardY)
			inRange = 0 <= newY && newY < len(data.Area) &&
				0 <= newX && newX < len(data.Area[newY])
			if !inRange || data.Area[newY][newX] != '#' {
				guardX = newX
				guardY = newY
				break
			}
			direction = turn(direction)
		}
	}
	return visited, true
}

func d6Part1(path string) {
	data := getD6Data(path)
	visited, _ := guardTravel(data)
	fmt.Printf("Visited spots: %d\n", len(visited))
}

func d6Part2(path string) {
	data := getD6Data(path)
	visited, _ := guardTravel(data)
	loops := 0
	for p, _ := range visited {
		r := data.Area[p.Y][p.X]
		data.Area[p.Y][p.X] = '#'
		if _, exited := guardTravel(data); !exited {
			loops += 1
		}
		data.Area[p.Y][p.X] = r
	}
	fmt.Printf("Loop potentials: %d\n", loops)
}

func Day6() {
	inputFile, err := filepath.Abs("./inputs/day6.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 6")
	d6Part1(inputFile)
	d6Part2(inputFile)
}
