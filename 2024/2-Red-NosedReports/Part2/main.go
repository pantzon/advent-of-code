package main

import (
	"fmt"
	"helpers"
	"strings"
)

func parser(line string) bool {
	if len(line) == 0 {
		return false
	}

	levels := strings.Split(line, " ")
	if checkLevels(levels) {
		return true
	}
	for i, _ := range levels {
		// Append has side-effects!
		newLevels := make([]string, len(levels))
		copy(newLevels, levels)
		newLevels = append(newLevels[:i], newLevels[i+1:]...)
		if checkLevels(newLevels) {
			return true
		}
	}
	return false
}

func checkLevels(levels []string) bool {
	lastLevel := levels[1]
	lastDirection, ok := checkDiff(levels[0], lastLevel)
	if !ok {
		return false
	}
	for _, level := range levels[2:] {
		direction, ok := checkDiff(lastLevel, level)
		if !ok || direction != lastDirection {
			return false
		}
		lastLevel = level
		lastDirection = direction
	}
	return true
}

func checkDiff(first, second string) (bool, bool) {
	diff := helpers.ParseInt(second) - helpers.ParseInt(first)
	if 0 < diff && diff < 4 {
		return true, true
	} else if -4 < diff && diff < 0 {
		return false, true
	}
	return false, false
}

func main() {
	flags := helpers.PrepCommonFlags()

	reports, err := helpers.ReadAndParseFile(helpers.ReadFileOptions[bool]{
		Path:   *(flags.Path),
		Parser: parser,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	safeReports := 0
	for _, safe := range reports {
		if safe {
			safeReports++
		}
	}

	fmt.Printf("Total safe reports: %d\n", safeReports)
}
