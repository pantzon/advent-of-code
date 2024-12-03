package main

import (
	"flag"
	"fmt"

	"aoc2024/days"
)

var dayFlag = flag.Int("day", 0, "")

var DayMap = map[int](func()){
	1:  days.Day1,
	2:  days.Day2,
	3:  days.Day3,
	4:  days.Day4,
	5:  days.Day5,
	6:  days.Day6,
	7:  days.Day7,
	8:  days.Day8,
	9:  days.Day9,
	10: days.Day10,
	11: days.Day11,
	12: days.Day12,
	13: days.Day13,
	14: days.Day14,
	15: days.Day15,
	16: days.Day16,
	17: days.Day17,
	18: days.Day18,
	19: days.Day19,
	20: days.Day20,
	21: days.Day21,
	22: days.Day22,
	23: days.Day23,
	24: days.Day24,
	25: days.Day25,
}

func main() {
	flag.Parse()

	if d, ok := DayMap[*dayFlag]; !ok {
		fmt.Printf("Unknown day: %d", *dayFlag)
	} else {
		d()
	}
}
