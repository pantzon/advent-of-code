package days

import (
	h "aoc/helpers"
	"fmt"
	"path/filepath"
)

type D16Node struct {
	Loc       h.Point
	Neighbors map[rune]*D16Node
}

type D16Map struct {
	Start *D16Node
	End   *D16Node
	Nodes map[h.Point]*D16Node
}

func d16Data(path string) D16Map {
	type accData struct {
		M     D16Map
		LastY int
	}
	data, err := h.ReduceFile(h.ReduceFileOptions[accData]{
		Path: path,
		Reducer: func(acc accData, line string) accData {
			acc.LastY += 1
			for i, r := range line {
				if r == 'S' || r == 'E' || r == '.' {
					p := h.Point{X: i, Y: acc.LastY}
					node := &D16Node{Loc: p, Neighbors: map[rune]*D16Node{}}
					acc.M.Nodes[p] = node
					if l, ok := acc.M.Nodes[p.Left()]; ok {
						node.Neighbors[DIR_LEFT] = l
						l.Neighbors[DIR_RIGHT] = node
					}
					if u, ok := acc.M.Nodes[p.Up()]; ok {
						node.Neighbors[DIR_UP] = u
						u.Neighbors[DIR_DOWN] = node
					}
					if r == 'S' {
						acc.M.Start = node
					}
					if r == 'E' {
						acc.M.End = node
					}
				}
			}
			return acc
		},
		InitialValue: accData{
			M:     D16Map{Nodes: map[h.Point]*D16Node{}},
			LastY: -1,
		},
	})
	if err != nil {
		panic(err)
	}
	// Remove dead ends
	// toRemove := map[*D16Node]bool{}
	// for _, node := range data.M.Nodes {
	// 	if node != data.M.Start && node != data.M.End && len(node.Neighbors) < 2 {
	// 		toRemove[node] = true
	// 	}
	// }
	// for len(toRemove) > 0 {
	// 	newRemove := map[*D16Node]bool{}
	// 	for node, _ := range toRemove {
	// 		delete(data.M.Nodes, node.Loc)
	// 		for dir, neighbor := range node.Neighbors {
	// 			deleteDir := reverseDir(dir)
	// 			delete(neighbor.Neighbors, deleteDir)
	// 			if neighbor != data.M.Start && neighbor != data.M.End && len(neighbor.Neighbors) < 2 {
	// 				newRemove[neighbor] = true
	// 			}
	// 		}
	// 	}
	// 	toRemove = newRemove
	// }
	return data.M
}

func getMover(dir rune) func(h.Point) h.Point {
	if dir == DIR_UP {
		return h.Point.Up
	}
	if dir == DIR_LEFT {
		return h.Point.Left
	}
	if dir == DIR_DOWN {
		return h.Point.Down
	}
	if dir == DIR_RIGHT {
		return h.Point.Right
	}
	fmt.Printf("Bad Direction (%s)!\n", string(dir))
	return func(p h.Point) h.Point { return p }
}

func reverseDir(dir rune) rune {
	if dir == DIR_UP {
		return DIR_DOWN
	}
	if dir == DIR_LEFT {
		return DIR_RIGHT
	}
	if dir == DIR_DOWN {
		return DIR_UP
	}
	if dir == DIR_RIGHT {
		return DIR_LEFT
	}
	fmt.Printf("Can't revertse! Bad Direction (%s)!\n", string(dir))
	return -1
}

func generateIndent(length int) string {
	i := 0
	out := make([]rune, length)
	for i < length {
		out[i] = ' '
		i++
	}
	return string(out)
}

type D16Visit struct {
	Dir    rune
	Turns  int
	Spaces int
	Path   []*D16Node
}

func (v D16Visit) weight() int {
	return v.Turns*1000 + v.Spaces
}

func cloneAndPrependPath(path []*D16Node, node *D16Node) []*D16Node {
	return append([]*D16Node{node}, path...)
}

func (m D16Map) Traverse(toVisit map[*D16Node][]D16Visit, visited map[*D16Node][]D16Visit) map[*D16Node][]D16Visit {
	if len(toVisit) == 0 {
		return visited
	}
	next := map[*D16Node][]D16Visit{}
	for orig, visits := range toVisit {
		visited[orig] = visits
		if orig == m.End {
			continue
		}
		for dir, neighbor := range orig.Neighbors {
			for _, visit := range visits {
				if dir != reverseDir(visit.Dir) {
					newVisit := D16Visit{Dir: dir, Turns: visit.Turns, Spaces: visit.Spaces + 1, Path: cloneAndPrependPath(visit.Path, neighbor)}
					if dir != visit.Dir {
						newVisit.Turns += 1
					}
					if prevs, ok := visited[neighbor]; !ok {
						next[neighbor] = append(next[neighbor], newVisit)
					} else {
						minPrev := prevs[0]
						for _, prev := range prevs[1:] {
							if minPrev.weight() > prev.weight() {
								minPrev = prev
							}
						}
						if minPrev.weight() > newVisit.weight() {
							delete(visited, neighbor)
							next[neighbor] = append(next[neighbor], newVisit)
						}
					}
				}
			}
		}
	}
	return m.Traverse(next, visited)
}

func d16Part1(path string) {
	m := d16Data(path)
	results := m.Traverse(map[*D16Node][]D16Visit{
		m.Start: {{Dir: DIR_RIGHT, Turns: 0, Spaces: 0, Path: []*D16Node{m.Start}}},
	}, map[*D16Node][]D16Visit{})
	minVisit := results[m.End][0]
	for _, visit := range results[m.End][1:] {
		if visit.weight() < minVisit.weight() {
			minVisit = visit
		}
	}
	fmt.Printf("Minimum: %v\n", minVisit.weight())
}

func d16Part2(path string) {
	m := d16Data(path)
	results := m.Traverse(map[*D16Node][]D16Visit{
		m.Start: {{Dir: DIR_RIGHT, Turns: 0, Spaces: 0, Path: []*D16Node{m.Start}}},
	}, map[*D16Node][]D16Visit{})
	minVisit := results[m.End][0]
	for _, visit := range results[m.End][1:] {
		if visit.weight() < minVisit.weight() {
			minVisit = visit
		}
	}
	uniqueSpots := map[*D16Node]bool{}
	for _, visit := range results[m.End] {
		if visit.weight() == minVisit.weight() {
			for _, node := range visit.Path {
				uniqueSpots[node] = true
			}
		}
	}
	fmt.Printf("Best Seats: %v\n", len(uniqueSpots))
}

func Day16() {
	inputFile, err := filepath.Abs("./inputs/day16.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 16")
	d16Part1(inputFile)
	d16Part2(inputFile)
}
