package days

import (
	h "aoc/helpers"
	"fmt"
	"path/filepath"
	"slices"
)

func MakeMemory(diskMap string) []int {
	var memory []int
	id := 0
	for i, s := range diskMap {
		size := h.ParseInt(string(s))
		val := -1
		if i%2 == 0 {
			val = id
			id++
		}
		memory = append(memory, slices.Repeat([]int{val}, size)...)
	}
	return memory
}

func P1ReworkMemory(memory []int) ([]int, int) {
	checksum := 0
	index := 0
	endIndex := len(memory) - 1
	for index < endIndex {
		for memory[endIndex] < 0 {
			endIndex--
		}
		if memory[index] < 0 {
			memory[index] = memory[endIndex]
			memory[endIndex] = -1
		}
		if memory[index] >= 0 {
			checksum += index * memory[index]
		}
		index++
	}
	for memory[index] >= 0 {
		checksum += index * memory[index]
		index++
	}
	return memory, checksum
}

func d9Part1(path string) {
	data, err := h.ReadFile(path)
	if err != nil {
		panic(err)
	}
	_, checksum := P1ReworkMemory(MakeMemory(data))
	fmt.Printf("Checksum: %d\n", checksum)
}

func P2ReworkMemory(memory []int) ([]int, int) {
	endIndex := len(memory) - 1
	for memory[endIndex] != 0 {
		for memory[endIndex] < 0 {
			endIndex--
		}
		size := endIndex
		id := memory[endIndex]
		for memory[endIndex] == id {
			endIndex--
		}
		size -= endIndex
		i := 0
		for i <= endIndex {
			for memory[i] >= 0 {
				i++
			}
			if i <= endIndex {
				startIndex := i
				space := i
				for memory[i] < 0 {
					i++
				}
				space = i - space
				if size <= space {
					j := 0
					for j < size {
						memory[startIndex+j] = id
						memory[endIndex+j+1] = -1
						j++
					}
					break
				}
			}
		}
	}
	checksum := 0
	for i, id := range memory {
		if id >= 0 {
			checksum += i * id
		}
	}
	return memory, checksum
}

func d9Part2(path string) {
	data, err := h.ReadFile(path)
	if err != nil {
		panic(err)
	}
	_, checksum := P2ReworkMemory(MakeMemory(data))
	fmt.Printf("Checksum: %d\n", checksum)
}

func Day9() {
	inputFile, err := filepath.Abs("./inputs/day9example.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 9")
	d9Part1(inputFile)
	d9Part2(inputFile)
}
