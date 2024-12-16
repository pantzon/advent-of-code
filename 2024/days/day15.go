package days

import (
	h "aoc/helpers"
	"fmt"
	"path/filepath"
)

const (
	DIR_UP    = '^'
	DIR_DOWN  = 'v'
	DIR_LEFT  = '<'
	DIR_RIGHT = '>'
)

type D15Map struct {
	Bounds h.Point
	Robot  h.Point
	Boxes  map[h.Point]bool
	Walls  map[h.Point]bool
}

type D15Data struct {
	Map          *D15Map
	Map2         *D15Map
	Instructions []rune
}

func d15Data(path string) D15Data {
	data, err := h.ReduceFile(h.ReduceFileOptions[D15Data]{
		Path: path,
		Reducer: func(acc D15Data, line string) D15Data {
			if len(line) > 0 {
				if line[0] == '#' {
					x := 0
					y := acc.Map.Bounds.Y
					for _, r := range line {
						p := h.Point{X: x, Y: y}
						p2 := h.Point{X: 2 * x, Y: y}
						p22 := h.Point{X: 2*x + 1, Y: y}
						if r == '#' {
							acc.Map.Walls[p] = true
							acc.Map2.Walls[p2] = true
							acc.Map2.Walls[p22] = true
						} else if r == 'O' {
							acc.Map.Boxes[p] = true
							acc.Map2.Boxes[p2] = true
						} else if r == '@' {
							acc.Map.Robot = p
							acc.Map2.Robot = p2
						}
						x += 1
					}
					acc.Map.Bounds = h.Point{X: x, Y: acc.Map.Bounds.Y + 1}
					acc.Map2.Bounds = h.Point{X: 2 * x, Y: acc.Map.Bounds.Y + 1}
				} else if line[0] == DIR_UP || line[0] == DIR_DOWN || line[0] == DIR_LEFT || line[0] == DIR_RIGHT {
					acc.Instructions = append(acc.Instructions, []rune(line)...)
				}
			}
			return acc
		},
		InitialValue: D15Data{
			Map: &D15Map{
				Boxes: map[h.Point]bool{},
				Walls: map[h.Point]bool{},
			},
			Map2: &D15Map{
				Boxes: map[h.Point]bool{},
				Walls: map[h.Point]bool{},
			},
		},
	})
	if err != nil {
		panic(err)
	}
	return data
}

func (m *D15Map) Move(inst rune) {
	var mover func(h.Point) h.Point
	if inst == DIR_UP {
		mover = h.Point.Up
	} else if inst == DIR_DOWN {
		mover = h.Point.Down
	} else if inst == DIR_LEFT {
		mover = h.Point.Left
	} else if inst == DIR_RIGHT {
		mover = h.Point.Right
	} else {
		fmt.Printf("Invalid Instruction (%s)\n", string(inst))
		return
	}
	m.applyMove(mover, m.Robot)
}

func (m *D15Map) applyMove(mover func(h.Point) h.Point, p h.Point) bool {
	isRobot := m.Robot == p
	moved := mover(p)
	if !m.Walls[moved] && (!m.Boxes[moved] || m.applyMove(mover, moved)) {
		if isRobot {
			m.Robot = moved
		} else {
			delete(m.Boxes, p)
			m.Boxes[moved] = true
		}
		return true
	}
	return false
}

func (m *D15Map) GPS() int {
	sum := 0
	for b, _ := range m.Boxes {
		sum += 100*b.Y + b.X
	}
	return sum
}

func (m *D15Map) Print() {
	y := 0
	for y < m.Bounds.Y {
		x := 0
		for x < m.Bounds.X {
			p := h.Point{X: x, Y: y}
			r := ' '
			if m.Robot == p {
				r = '@'
			} else if m.Boxes[p] {
				r = 'O'
			} else if m.Walls[p] {
				r = '#'
			}
			fmt.Printf("%s", string(r))
			x++
		}
		fmt.Println()
		y++
	}
}

func d15Part1(path string) {
	d := d15Data(path)
	// d.Map.Print()
	for _, r := range d.Instructions {
		d.Map.Move(r)
		// d.Map.Print()
	}
	fmt.Printf("GPS Sum: %d\n", d.Map.GPS())
}

func (m *D15Map) Move2(inst rune) {
	p := m.Robot
	if inst == DIR_LEFT {
		p = p.Left()
		for p.X > 0 {
			if m.Boxes[p] || m.Boxes[p.Left()] {
				p = p.Left()
			} else {
				break
			}
		}
		if !m.Walls[p] {
			for p.X < m.Robot.X {
				if m.Boxes[p] {
					delete(m.Boxes, p)
					m.Boxes[p.Left()] = true
					p = p.Right()
				}
				p = p.Right()
			}
			m.Robot = p.Left()
		}
	} else if inst == DIR_RIGHT {
		p = p.Right()
		for p.X < m.Bounds.X {
			if m.Boxes[p] {
				p = p.Right().Right()
			} else {
				break
			}
		}
		if !m.Walls[p] {
			for p.X > m.Robot.X {
				if m.Boxes[p] {
					delete(m.Boxes, p)
					m.Boxes[p.Right()] = true
				}
				p = p.Left()
			}
			m.Robot = p.Right()
		}
	} else {
		mover := h.Point.Up
		if inst == DIR_DOWN {
			mover = h.Point.Down
		}
		m.applyMoveVert(mover, []h.Point{m.Robot})
	}
}

func (m *D15Map) applyMoveVert(mover func(h.Point) h.Point, toCheck []h.Point) bool {
	newChecks := map[h.Point]bool{}
	anyBoxes := false
	for _, p := range toCheck {
		moved := mover(p)
		if m.Walls[moved] {
			return false
		}
		if m.Boxes[moved] {
			newChecks[moved] = true
			newChecks[moved.Right()] = true
			anyBoxes = true
		} else if m.Boxes[moved.Left()] {
			newChecks[moved] = true
			newChecks[moved.Left()] = true
			anyBoxes = true
		}
	}
	var keys []h.Point
	for p, _ := range newChecks {
		keys = append(keys, p)
	}
	if !anyBoxes || m.applyMoveVert(mover, keys) {
		if len(toCheck) == 1 && toCheck[0] == m.Robot {
			m.Robot = mover(toCheck[0])
		} else {
			for _, p := range toCheck {
				if m.Boxes[p] {
					delete(m.Boxes, p)
					m.Boxes[mover(p)] = true
				}
			}
		}
		return true
	}
	return false
}

func (m *D15Map) Print2() {
	y := 0
	for y < m.Bounds.Y {
		x := 0
		for x < m.Bounds.X {
			p := h.Point{X: x, Y: y}
			r := ' '
			if m.Robot == p {
				r = '@'
			} else if m.Boxes[p] {
				r = '['
			} else if m.Boxes[h.Point{X: p.X - 1, Y: p.Y}] {
				r = ']'
			} else if m.Walls[p] {
				r = '#'
			}
			fmt.Printf("%s", string(r))
			x++
		}
		fmt.Println()
		y++
	}
}

func d15Part2(path string) {
	d := d15Data(path)
	// d.Map2.Print2()
	for _, r := range d.Instructions {
		d.Map2.Move2(r)
		// d.Map2.Print2()
	}
	fmt.Printf("GPS Sum: %d\n", d.Map2.GPS())
}

func Day15() {
	inputFile, err := filepath.Abs("./inputs/day15.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 15")
	d15Part1(inputFile)
	d15Part2(inputFile)
}
