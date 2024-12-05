package days

import (
	"container/heap"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"aoc/helpers"
)

func d1Part1(path string) {
	left := &helpers.Heap[int]{}
	heap.Init(left)
	right := &helpers.Heap[int]{}
	heap.Init(right)

	parser := func(line string) string {
		if len(line) == 0 {
			return ""
		}
		pieces := strings.Split(line, " ")
		first := pieces[0]
		last := pieces[len(pieces)-1]
		if val, err := strconv.Atoi(first); err == nil {
			heap.Push(left, val)
		} else {
			fmt.Printf("Error %s for (left) parsing of %s\n", err, first)
			fmt.Println(err)
		}
		if val, err := strconv.Atoi(last); err == nil {
			heap.Push(right, val)
		} else {
			fmt.Printf("Error %s for (right) parsing of %s\n", err, last)
			fmt.Println(err)
		}
		return ""
	}

	if _, err := helpers.ParseFile(helpers.ParseFileOptions[string]{
		Path:   path,
		Parser: parser,
	}); err != nil {
		fmt.Println(err)
		return
	}

	total := 0
	for left.Len() > 0 && right.Len() > 0 {
		diff := heap.Pop(left).(int) - heap.Pop(right).(int)
		if diff < 0 {
			diff = -1 * diff
		}
		total += diff
	}

	fmt.Printf("Total diff distance: %d\n", total)
}

func d1Part2(path string) {
	left := map[string]int{}
	right := map[string]int{}

	parser := func(line string) string {
		if len(line) == 0 {
			return ""
		}
		pieces := strings.Split(line, " ")
		first := pieces[0]
		last := pieces[len(pieces)-1]
		if _, ok := left[first]; !ok {
			left[first] = 0
		}
		left[first]++
		if _, ok := right[last]; !ok {
			right[last] = 0
		}
		right[last]++
		return ""
	}

	if _, err := helpers.ParseFile(helpers.ParseFileOptions[string]{
		Path:   path,
		Parser: parser,
	}); err != nil {
		fmt.Println(err)
		return
	}

	total := 0
	for key, count := range left {
		rCount, ok := right[key]
		if !ok {
			rCount = 0
		}
		keyNum, err := strconv.Atoi(key)
		if err != nil {
			fmt.Println(err)
			return
		}
		total += keyNum * count * rCount
	}

	fmt.Printf("Total similarity: %d\n", total)
}

func Day1() {
	inputFile, err := filepath.Abs("./inputs/day1.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 1")
	d1Part1(inputFile)
	d1Part2(inputFile)
}
