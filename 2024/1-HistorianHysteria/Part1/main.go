package main

import (
	"container/heap"
	"flag"
	"fmt"
	"helpers"
	"strconv"
	"strings"
)

func main() {
	pathPtr := flag.String("path", "", "")
	flag.Parse()

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

	if _, err := helpers.ReadAndParseFile(helpers.ReadFileOptions[string]{
		Path:   *pathPtr,
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
