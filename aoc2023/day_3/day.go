package day_3

import (
	"advent-of-code/aoc"
	"advent-of-code/aoc/utils"
	"log"
	"slices"
	"strconv"
	"text/scanner"
)

type day struct {
	inputFile string
}

var _ aoc.EventRunner = (*day)(nil)

func Create() aoc.EventRunner {
	return create("aoc2023/day_3/input.txt")
}

func create(inputFile string) *day {
	return &day{
		inputFile: inputFile,
	}
}

func (d *day) Execute() {
	log.Println("Day 3")
	log.Println("Q1:", d.Question1())
	log.Println("Q2:", d.Question2())
}

func (d *day) Question1() string {
	n, s, _ := d.loadData()
	sum := 0
	for _, v := range listOfValidNumbers(n, s) {
		sum += v
	}

	return strconv.Itoa(sum)
}

func listOfValidNumbers(numbers []number, symbols map[int][]int) []int {
	res := make([]int, 0)
	for _, n := range numbers {
		for i := n.pos.Column; i < n.pos.Column+len(n.stringValue); i++ {
			if isSymbolAdjacent(symbols[n.pos.Line-1], i) || isSymbolAdjacent(symbols[n.pos.Line], i) || isSymbolAdjacent(symbols[n.pos.Line+1], i) {
				res = append(res, n.value)
				break
			}
		}
	}

	return res
}

func isSymbolAdjacent(symbols []int, x int) bool {
	if len(symbols) == 0 {
		return false
	}

	return slices.Contains(symbols, x-1) || slices.Contains(symbols, x) || slices.Contains(symbols, x+1)
}

func isSymbolAdjacentWithIndex(symbols []int, x int) (bool, int) {
	if len(symbols) == 0 {
		return false, 0
	}

	if slices.Contains(symbols, x-1) {
		return true, x - 1
	}

	if slices.Contains(symbols, x) {
		return true, x
	}

	if slices.Contains(symbols, x+1) {
		return true, x + 1
	}
	return false, 0
}

func (d *day) Question2() string {
	n, _, g := d.loadData()

	sum := 0
	for _, v := range listOfValidGears(n, g) {
		for _, v1 := range v {

			if len(v1) == 2 {
				sum = sum + (v1[0] * v1[1])
			}
		}
	}

	return strconv.Itoa(sum)
}

func listOfValidGears(numbers []number, gears map[int][]int) map[int]map[int][]int {
	res := make(map[int]map[int][]int)
	for _, n := range numbers {
		for i := n.pos.Column; i < n.pos.Column+len(n.stringValue); i++ {
			if ok, ind := isSymbolAdjacentWithIndex(gears[n.pos.Line-1], i); ok {
				if _, ok := res[n.pos.Line-1]; !ok {
					res[n.pos.Line-1] = make(map[int][]int)
				}
				res[n.pos.Line-1][ind] = append(res[n.pos.Line-1][ind], n.value)
				break
			}
			if ok, ind := isSymbolAdjacentWithIndex(gears[n.pos.Line], i); ok {
				if _, ok := res[n.pos.Line]; !ok {
					res[n.pos.Line] = make(map[int][]int)
				}
				res[n.pos.Line][ind] = append(res[n.pos.Line][ind], n.value)
				break
			}
			if ok, ind := isSymbolAdjacentWithIndex(gears[n.pos.Line+1], i); ok {
				if _, ok := res[n.pos.Line+1]; !ok {
					res[n.pos.Line+1] = make(map[int][]int)
				}
				res[n.pos.Line+1][ind] = append(res[n.pos.Line+1][ind], n.value)
				break
			}
		}
	}

	return res
}

type number struct {
	stringValue string
	pos         scanner.Position
	value       int
}

func (d *day) loadData() ([]number, map[int][]int, map[int][]int) {
	file := utils.LoadInput(d.inputFile)

	symbols := make(map[int][]int)
	numbers := make([]number, 0)
	gears := make(map[int][]int)
	var s scanner.Scanner
	s.Init(file)
	s.Mode = scanner.ScanInts

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		if s.TokenText() == "." {
			continue
		}
		if v, err := strconv.Atoi(s.TokenText()); err == nil {
			numbers = append(numbers, number{stringValue: s.TokenText(), value: v, pos: s.Position})
		} else {
			symbols[s.Line] = append(symbols[s.Line], s.Column)
			if s.TokenText() == "*" {
				gears[s.Line] = append(gears[s.Line], s.Column)
			}
		}
	}

	return numbers, symbols, gears
}
