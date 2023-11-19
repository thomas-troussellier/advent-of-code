package day_3

import (
	"advent-of-code/aoc"
	"advent-of-code/aoc/utils"
	"bufio"
	"log"
	"strconv"
	"strings"
)

type day struct {
	inputFile string
}

var _ aoc.EventRunner = (*day)(nil)

func Create() aoc.EventRunner {
	return create("aoc2022/day_3/input.txt")
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
	rumsacks := loadRumsackData(d.inputFile)

	return strconv.Itoa(findPriorities1(rumsacks))
}

func (d *day) Question2() string {
	rumsacks := loadRumsackData(d.inputFile)

	return strconv.Itoa(findPriorities2(rumsacks))
}

func findPriorities1(rumsacks []string) int {
	totalPriorities := 0
	for _, rumsack := range rumsacks {
		// split into compartments
		half := len(rumsack) / 2
		part1 := rumsack[:half]
		part2 := rumsack[half:]

		// find common letter in both packs
		//var part1Index, part2Index int
		var found rune
		for _, c := range part1 {
			if index := strings.IndexRune(part2, c); index >= 0 {
				//part1Index = i
				//part2Index = index
				found = c
				break
			}
		}

		var priority int
		if found > 91 {
			priority = int(found) - 96
		} else {
			priority = int(found) - 38
		}

		totalPriorities += priority
	}
	return totalPriorities
}

func findPriorities2(rumsacks []string) int {
	totalPriorities := 0

	groups := make([][]string, 0)
	for i := 0; i <= len(rumsacks)-3; i += 3 {
		list := rumsacks[i : i+3]
		groups = append(groups, list)
	}

	for _, list := range groups {

		// find common letter in all lists
		var found rune
		for _, c := range list[0] {
			i1 := strings.IndexRune(list[1], c)
			i2 := strings.IndexRune(list[2], c)
			if i1 >= 0 && i2 >= 0 {
				found = c
				break
			}
		}

		var priority int
		if found > 91 {
			priority = int(found) - 96
		} else {
			priority = int(found) - 38
		}

		totalPriorities += priority
	}
	return totalPriorities
}

func loadRumsackData(fileName string) []string {
	readFile := utils.LoadInput(fileName)

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	rumsack := make([]string, 0)

	for fileScanner.Scan() {
		currentLine := fileScanner.Text()

		rumsack = append(rumsack, currentLine)
	}
	readFile.Close()

	return rumsack
}
