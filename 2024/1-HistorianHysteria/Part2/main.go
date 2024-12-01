package main

import (
	"flag"
	"fmt"
	"helpers/io"
	"strconv"
	"strings"
)

func main() {
	pathPtr := flag.String("path", "", "")
	flag.Parse()

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

	if _, err := io.ReadAndParseFile(io.ReadFileOptions[string]{
		Path:   *pathPtr,
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
