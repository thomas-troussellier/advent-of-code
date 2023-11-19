package day_8

import (
	"advent-of-code/aoc"
	"advent-of-code/aoc/utils"
	"bufio"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

type day struct {
	inputFile string
}

var _ aoc.EventRunner = (*day)(nil)

func Create() aoc.EventRunner {
	return create("aoc2015/day_8/input.txt")
}

func create(inputFile string) *day {
	return &day{
		inputFile: inputFile,
	}
}

func (d *day) Execute() {
	log.Println("Day 8")
	log.Println("Q1:", d.Question1())
	log.Println("Q2:", d.Question2())
}

func (d *day) Question1() string {
	matchsticks := loadMatchsticks(d.inputFile)

	return strconv.Itoa(parseData(matchsticks))
}

func (d *day) Question2() string {
	matchsticks := loadMatchsticks(d.inputFile)

	return strconv.Itoa(parseData2(matchsticks))
}

func parseData(matchsticks []string) int {
	charTotal := 0
	memTotal := 0

	for _, s := range matchsticks {
		charTotal += len(s)
		memTotal += memSize(s)
	}

	return charTotal - memTotal
}

func parseData2(matchsticks []string) int {
	charTotal := 0
	newCode := 0

	// add \" around each s[1:len(s)-1]
	// add \ on each \[x"]

	for _, s := range matchsticks {
		toAnalyse := replaceable.ReplaceAllStringFunc(s[1:len(s)-1], func(input string) string {
			if input == "\\\"" || input == "\\\\" {
				return fmt.Sprintf("\\\\%s", input)
			} else {
				return fmt.Sprintf("\\%s", input)
			}
		})
		charTotal += len(s)
		newCode = newCode + len(toAnalyse) + 6
	}

	return newCode - charTotal
}

var replaceable = regexp.MustCompile(`\\"|\\\\|\\x[0-9A-Fa-f]{2}`)

func memSize(s string) int {
	// remove " at start & end
	toAnalyse := s[1 : len(s)-1]

	return len(replaceable.ReplaceAllString(toAnalyse, "X"))
}

func loadMatchsticks(fileName string) []string {
	readFile := utils.LoadInput(fileName)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	list := make([]string, 0)

	for fileScanner.Scan() {
		instruction := fileScanner.Text()

		list = append(list, instruction)
	}

	readFile.Close()

	return list
}
