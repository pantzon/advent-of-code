package days

import (
	"aoc/helpers"
	"fmt"
	"path/filepath"
	"strings"
)

type EquationBase struct {
	Result   int
	Operands []int
}

func parser(val string) EquationBase {
	pieces := strings.Split(val, " ")
	eb := EquationBase{
		Result: helpers.ParseInt(strings.TrimSuffix(pieces[0], ":")),
	}
	for _, num := range pieces[1:] {
		eb.Operands = append(eb.Operands, helpers.ParseInt(num))
	}
	return eb
}

func EbTestPart1(result int, pieces []int) bool {
	if len(pieces) == 1 {
		return result == pieces[0]
	}
	return EbTestPart1(result, append([]int{pieces[0] + pieces[1]}, pieces[2:]...)) ||
		EbTestPart1(result, append([]int{pieces[0] * pieces[1]}, pieces[2:]...))
}

func d7Part1(path string) {
	data, err := helpers.ParseFile(helpers.ParseFileOptions[EquationBase]{
		Path:   path,
		Parser: parser,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	total := 0
	for _, eb := range data {
		if EbTestPart1(eb.Result, eb.Operands) {
			total += eb.Result
		}
	}
	fmt.Printf("Total calibration result: %d\n", total)
}

func EbTestPart2(result int, pieces []int) bool {
	if len(pieces) == 1 {
		return result == pieces[0]
	}
	return EbTestPart2(result, append([]int{pieces[0] + pieces[1]}, pieces[2:]...)) ||
		EbTestPart2(result, append([]int{pieces[0] * pieces[1]}, pieces[2:]...)) ||
		EbTestPart2(result, append([]int{helpers.ParseInt(fmt.Sprintf("%d%d", pieces[0], pieces[1]))}, pieces[2:]...))
}

func d7Part2(path string) {
	data, err := helpers.ParseFile(helpers.ParseFileOptions[EquationBase]{
		Path:   path,
		Parser: parser,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	total := 0
	for _, eb := range data {
		if EbTestPart2(eb.Result, eb.Operands) {
			total += eb.Result
		}
	}
	fmt.Printf("Total calibration result: %d\n", total)
}

func Day7() {
	inputFile, err := filepath.Abs("./inputs/day7.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 7")
	d7Part1(inputFile)
	d7Part2(inputFile)
}
