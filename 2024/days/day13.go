package days

import (
	h "aoc/helpers"
	"fmt"
	"path/filepath"
	"regexp"
)

type ClawGame struct {
	A     h.Point
	B     h.Point
	Prize h.Point
}

var (
	A_REGEX     = regexp.MustCompile("Button A: X\\+([0-9]+), Y\\+([0-9]+)")
	B_REGEX     = regexp.MustCompile("Button B: X\\+([0-9]+), Y\\+([0-9]+)")
	PRIZE_REGEX = regexp.MustCompile("Prize: X=([0-9]+), Y=([0-9]+)")
)

func d13Data(path string) []*ClawGame {
	type ReducerData struct {
		Games    []*ClawGame
		CurrGame *ClawGame
	}
	data, err := h.ReduceFile(h.ReduceFileOptions[ReducerData]{
		Path: path,
		Reducer: func(acc ReducerData, line string) ReducerData {
			if acc.CurrGame == nil {
				acc.CurrGame = &ClawGame{}
				acc.Games = append(acc.Games, acc.CurrGame)
			}
			if m := A_REGEX.FindSubmatch([]byte(line)); m != nil {
				acc.CurrGame.A = h.Point{X: h.ParseInt(string(m[1])), Y: h.ParseInt(string(m[2]))}
			} else if m = B_REGEX.FindSubmatch([]byte(line)); m != nil {
				acc.CurrGame.B = h.Point{X: h.ParseInt(string(m[1])), Y: h.ParseInt(string(m[2]))}
			} else if m = PRIZE_REGEX.FindSubmatch([]byte(line)); m != nil {
				acc.CurrGame.Prize = h.Point{X: h.ParseInt(string(m[1])), Y: h.ParseInt(string(m[2]))}
				acc.CurrGame = nil
			}
			return acc
		},
		InitialValue: ReducerData{},
	})
	if err != nil {
		panic(err)
	}
	return data.Games
}

func PrizeCost(game *ClawGame, limit int) (int, bool) {
	denominator := game.A.X*game.B.Y - game.A.Y*game.B.X
	bNumerator := game.Prize.Y*game.A.X - game.Prize.X*game.A.Y
	if bNumerator%denominator == 0 {
		aNumerator := game.Prize.Y*game.B.X - game.Prize.X*game.B.Y
		if aNumerator%denominator == 0 {
			if limit == 0 || (limit >= -1*(aNumerator/denominator) && limit >= (bNumerator/denominator)) {
				winner := (aNumerator/denominator)*-3 + (bNumerator / denominator)
				return winner, true
			}
		}
	}
	return 0, false
}

func d13Part1(path string) {
	games := d13Data(path)
	total := 0
	for _, game := range games {
		if cost, winnable := PrizeCost(game, 100); winnable {
			total += cost
		}
	}
	fmt.Printf("Total Cost: %d\n", total)
}

func d13Part2(path string) {
	games := d13Data(path)
	total := 0
	for _, game := range games {
		game.Prize = h.Point{X: game.Prize.X + 10000000000000, Y: game.Prize.Y + 10000000000000}
		if cost, winnable := PrizeCost(game, 0); winnable {
			total += cost
		}
	}
	fmt.Printf("Total Cost: %d\n", total)
}

func Day13() {
	inputFile, err := filepath.Abs("./inputs/day13.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 13")
	d13Part1(inputFile)
	d13Part2(inputFile)
}
