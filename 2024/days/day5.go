package days

import (
	"aoc/helpers"
	"fmt"
	"path/filepath"
	"slices"
	"strings"
)

type Page struct {
	Id       string
	Parents  map[string]*Page
	Children map[string]*Page
}

type Data struct {
	Pages     map[string]*Page
	PrintSets [][]*Page
}

func getOrMakePage(data Data, id string) *Page {
	if _, ok := data.Pages[id]; !ok {
		data.Pages[id] = &Page{
			Id:       id,
			Parents:  map[string]*Page{},
			Children: map[string]*Page{},
		}
	}
	return data.Pages[id]
}

func getData(path string) *Data {
	reducer := func(acc Data, line string) Data {
		if strings.Contains(line, "|") {
			pieces := strings.Split(line, "|")
			if len(pieces) != 2 {
				fmt.Printf("Invalid dependency mapping! %s\n", line)
				return acc
			}
			first := getOrMakePage(acc, pieces[0])
			second := getOrMakePage(acc, pieces[1])
			first.Children[second.Id] = second
			second.Parents[first.Id] = first
		} else if strings.Contains(line, ",") {
			var pages []*Page
			for _, id := range strings.Split(line, ",") {
				pages = append(pages, getOrMakePage(acc, id))
			}
			if len(pages)%2 != 1 {
				fmt.Printf("Even number of pages in print set! %s\n", line)
				return acc
			}
			acc.PrintSets = append(acc.PrintSets, pages)
		}
		return acc
	}
	data, err := helpers.ReduceFile(helpers.ReduceFileOptions[Data]{
		Path:    path,
		Reducer: reducer,
		InitialValue: Data{
			Pages: map[string]*Page{},
		},
	})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &data
}

func comparePages(a, b *Page) int {
	if _, ok := a.Children[b.Id]; ok {
		return -1
	} else if _, ok := a.Parents[b.Id]; ok {
		return 1
	} else {
		for _, child := range a.Children {
			comp := comparePages(child, b)
			if comp < 0 {
				return comp
			}
		}
		for _, parent := range a.Parents {
			comp := comparePages(parent, b)
			if comp > 0 {
				return comp
			}
		}
	}
	return 0
}

func splitPrintSets(data *Data) ([][]*Page, [][]*Page) {
	var valids [][]*Page
	var invalids [][]*Page

	for _, printSet := range data.PrintSets {
		if slices.IsSortedFunc(printSet, comparePages) {
			valids = append(valids, printSet)
		} else {
			invalids = append(invalids, printSet)
		}
	}
	return valids, invalids
}

func d5Part1(path string) {
	valids, _ := splitPrintSets(getData(path))

	midSums := 0
	for _, printSet := range valids {
		midSums += helpers.ParseInt(printSet[(len(printSet)-1)/2].Id)
	}
	fmt.Printf("Sum of Valid Middles: %d\n", midSums)
}

func d5Part2(path string) {
	data := getData(path)
	_, invalids := splitPrintSets(data)

	midSums := 0
	for _, printSet := range invalids {
		slices.SortFunc(printSet, comparePages)
		midSums += helpers.ParseInt(printSet[(len(printSet)-1)/2].Id)
	}
	fmt.Printf("Sum of Fixed Middles: %d\n", midSums)
}

func Day5() {
	inputFile, err := filepath.Abs("./inputs/day5.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 5")
	d5Part1(inputFile)
	d5Part2(inputFile)
}
