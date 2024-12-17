package days

import (
	h "aoc/helpers"
	"fmt"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
)

type D17Computer struct {
	A           int
	B           int
	C           int
	Program     []int
	InstPointer int
	Output      []int
}

var operations = map[int]func(c *D17Computer, operand int) int{
	0: (*D17Computer).adv,
	1: (*D17Computer).bxl,
	2: (*D17Computer).bst,
	3: (*D17Computer).jnz,
	4: (*D17Computer).bxc,
	5: (*D17Computer).out,
	6: (*D17Computer).bdv,
	7: (*D17Computer).cdv,
}

func (c *D17Computer) IsActive() bool {
	return c.IsValid() && c.InstPointer < len(c.Program)
}

func (c *D17Computer) IsValid() bool {
	return c.InstPointer >= 0
}

func (c *D17Computer) Reset(a int) {
	c.A = a
	c.B = 0
	c.C = 0
	c.InstPointer = 0
	c.Output = nil
}

func (c *D17Computer) Debug() {
	fmt.Println(">>> COMPUTER STATUS",
		"\n>>>  A: ", c.A,
		"\n>>>  B: ", c.B,
		"\n>>>  C: ", c.C,
		"\n>>>  Inst: ", c.InstPointer)
	fmt.Println("")
}

func (c *D17Computer) Print() {
	out := make([]string, len(c.Output))
	i := 0
	for i < len(out) {
		out[i] = fmt.Sprintf("%d", c.Output[i])
		i++
	}
	fmt.Println("Output: ", strings.Join(out, ","))
}

func (c *D17Computer) RunProgram() bool {
	// c.Debug()
	for c.ExecuteNext() {
		// c.Debug()
	}
	return true
}

func (c *D17Computer) ExecuteNext() bool {
	c.InstPointer = operations[c.Program[c.InstPointer]](c, c.Program[c.InstPointer+1])
	return c.IsActive()
}

func (c *D17Computer) adv(operand int) int {
	c.A = c.A >> c.combo(operand)
	return c.InstPointer + 2
}

func (c *D17Computer) bxl(operand int) int {
	c.B = c.B ^ operand
	return c.InstPointer + 2
}

func (c *D17Computer) bst(operand int) int {
	c.B = c.combo(operand) % 8
	return c.InstPointer + 2
}

func (c *D17Computer) jnz(operand int) int {
	if c.A != 0 {
		return operand
	}
	return c.InstPointer + 2
}

func (c *D17Computer) bxc(_ int) int {
	c.B = c.B ^ c.C
	return c.InstPointer + 2
}

func (c *D17Computer) out(operand int) int {
	c.Output = append(c.Output, c.combo(operand)%8)
	return c.InstPointer + 2
}

func (c *D17Computer) bdv(operand int) int {
	c.B = c.A >> c.combo(operand)
	return c.InstPointer + 2
}

func (c *D17Computer) cdv(operand int) int {
	c.C = c.A >> c.combo(operand)
	return c.InstPointer + 2
}

func (c *D17Computer) combo(operand int) int {
	if operand == 4 {
		return c.A
	}
	if operand == 5 {
		return c.B
	}
	if operand == 6 {
		return c.C
	}
	if operand < 0 || operand > 6 {
		fmt.Println("BAD OPERAND! ", operand)
	}
	return operand
}

var REGISTER_RE = regexp.MustCompile("^Register ([ABC]): ([0-9]+)$")
var PROGRAM_RE = regexp.MustCompile("^Program: ([0-7,]+)$")

func d17Data(path string) *D17Computer {
	data, err := h.ReduceFile(h.ReduceFileOptions[*D17Computer]{
		Path: path,
		Reducer: func(acc *D17Computer, line string) *D17Computer {
			if matches := REGISTER_RE.FindSubmatch([]byte(line)); matches != nil {
				if "A" == string(matches[1]) {
					acc.A = h.ParseInt(string(matches[2]))
				} else if "B" == string(matches[1]) {
					acc.B = h.ParseInt(string(matches[2]))
				} else if "C" == string(matches[1]) {
					acc.C = h.ParseInt(string(matches[2]))
				}
			}
			if matches := PROGRAM_RE.FindSubmatch([]byte(line)); matches != nil {
				for _, i := range strings.Split(string(matches[1]), ",") {
					acc.Program = append(acc.Program, h.ParseInt(i))
				}
			}
			return acc
		},
		InitialValue: &D17Computer{},
	})
	if err != nil {
		panic(err)
	}
	return data
}

func d17Part1(path string) {
	computer := d17Data(path)
	computer.RunProgram()
	computer.Print()
}

func (c *D17Computer) findProgramMatch(a int, want []int) []int {
	var matches []int
	i := a
	// 8 is a magic number, as each iteration A = A/8 (or A = A >> 3), meaning that the
	// next iteration only really depends on (A << 3 + last 3 bits), which is 8 possibilities.
	for i < a+8 {
		c.Reset(i)
		c.RunProgram()
		if slices.Equal(c.Output, want) {
			if len(want) == len(c.Program) {
				matches = append(matches, i)
			} else {
				matches = append(matches, c.findProgramMatch(i<<3, c.Program[len(c.Program)-len(want)-1:])...)
			}
		}
		i++
	}
	return matches
}

func d17Part2(path string) {
	computer := d17Data(path)
	minMatch := -1
	for _, m := range computer.findProgramMatch(0, computer.Program[len(computer.Program)-1:]) {
		if minMatch < 0 || m < minMatch {
			minMatch = m
		}
	}
	fmt.Println("New A: ", minMatch)
}

func Day17() {
	inputFile, err := filepath.Abs("./inputs/day17.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 17")
	d17Part1(inputFile)
	d17Part2(inputFile)
}
