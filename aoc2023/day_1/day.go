package day_1

import (
	"advent-of-code/aoc"
	"advent-of-code/aoc/utils"
	"bufio"
	"log"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

type day struct {
	inputFile string
}

var _ aoc.EventRunner = (*day)(nil)

func Create() aoc.EventRunner {
	return create("aoc2023/day_1/input.txt")
}

func create(inputFile string) *day {
	return &day{
		inputFile: inputFile,
	}
}

func (d *day) Execute() {
	log.Println("Day 1")
	log.Println("Q1:", d.Question1())
	log.Println("Q2:", d.Question2())
}

func (d *day) Question1() string {
	data := d.loadData()

	sum := 0
	for _, s := range data {
		sum += parseNumber(s, false)
	}
	return strconv.Itoa(sum)
}

func (d *day) Question2() string {
	data := d.loadData()

	sum := 0
	for _, s := range data {
		sum += parseNumber(s, true)
	}
	return strconv.Itoa(sum)
}

func parseNumber(s string, useAlpha bool) int {
	var (
		val    string = ""
		resTab []t
	)
	fIndex := strings.IndexFunc(s, unicode.IsDigit)
	lIndex := strings.LastIndexFunc(s, unicode.IsDigit)
	if useAlpha {
		tab := alphaDigitOccurences(s)
		if len(tab) > 0 {
			resTab = append(resTab, tab[0])
			if len(tab) > 1 {
				resTab = append(resTab, tab[len(tab)-1])
			}
		}
	}

	if fIndex >= 0 {
		resTab = append(resTab, t{value: string(s[fIndex]), index: fIndex})
	}
	if lIndex >= 0 {
		resTab = append(resTab, t{value: string(s[lIndex]), index: lIndex})
	}

	if len(resTab) > 0 {
		slices.SortFunc(resTab, func(a, b t) int {
			return a.index - b.index
		})

		val = resTab[0].value
		if len(resTab) > 1 {
			val += resTab[len(resTab)-1].value
		}
	}

	ival, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal("failed", err)
	}

	return ival
}

var alphaDigit = map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}

type t struct {
	value string
	index int
}

func (v t) String() string {
	return strconv.Itoa(v.index)
}

func alphaDigitOccurences(input string) []t {
	res := make([]t, 0)
	for digit, value := range alphaDigit {
		digit := digit
		value := value
		if count := strings.Count(input, digit); count > 0 {
			res = append(res, t{value: value, index: strings.Index(input, digit)})
			// if only one occurence, we actually want to have it as both start and finish
			res = append(res, t{value: value, index: strings.LastIndex(input, digit)})
		}
	}
	if len(res) == 0 {
		return res
	}

	slices.SortFunc(res, func(a, b t) int {
		return a.index - b.index
	})
	return res
}

func (d *day) loadData() []string {
	file := utils.LoadInput(d.inputFile)

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	res := make([]string, 0)
	for fileScanner.Scan() {
		res = append(res, fileScanner.Text())
	}

	return res
}
