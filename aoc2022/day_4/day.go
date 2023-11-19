package day_4

import (
	"advent-of-code/aoc"
	"advent-of-code/aoc/utils"
	"bufio"
	"fmt"
	"log"
	"strconv"
)

type day struct {
	inputFile string
}

var _ aoc.EventRunner = (*day)(nil)

func Create() aoc.EventRunner {
	return create("aoc2022/day_4/input.txt")
}

func create(inputFile string) *day {
	return &day{
		inputFile: inputFile,
	}
}

func (d *day) Execute() {
	log.Println("Day 4")
	log.Println("Q1:", d.Question1())
	log.Println("Q2:", d.Question2())
}

func (d *day) Question1() string {
	pairsData := loadSpacePairsData(d.inputFile)

	return strconv.Itoa(findContained(pairsData))
}

func (d *day) Question2() string {
	pairsData := loadSpacePairsData(d.inputFile)

	return strconv.Itoa(findOverlapped(pairsData))
}

func findOverlapped(pairsData [][]spaceRange) int {
	overlapped := 0

	for _, pairs := range pairsData {
		if checkOverlap(pairs[0], pairs[1]) {
			overlapped++
		}
	}

	return overlapped
}

func checkOverlap(p1, p2 spaceRange) bool {
	if p1.start <= p2.start && p1.end >= p2.end {
		return true
	}
	if p1.start <= p2.start && p1.end >= p2.start && p1.end <= p2.end {
		return true
	}
	if p2.start <= p1.start && p2.end >= p1.start && p1.end >= p2.end {
		return true
	}
	if p2.start <= p1.start && p1.end <= p2.end {
		return true
	}

	return false
}

func findContained(pairsData [][]spaceRange) int {
	contained := 0

	for _, pairs := range pairsData {
		if checkContained(pairs[0], pairs[1]) {
			contained++
		}
	}

	return contained
}

func checkContained(p1, p2 spaceRange) bool {
	if p1.start <= p2.start && p1.end >= p2.end {
		return true
	}
	if p2.start <= p1.start && p1.end <= p2.end {
		return true
	}

	return false
}

func loadSpacePairsData(fileName string) [][]spaceRange {
	readFile := utils.LoadInput(fileName)
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	cleaningAreaPairs := make([][]spaceRange, 0)

	for fileScanner.Scan() {
		currentLine := fileScanner.Text()

		var sP1, sP2, eP1, eP2 int
		fmt.Sscanf(currentLine, "%d-%d,%d-%d", &sP1, &eP1, &sP2, &eP2)

		pair1 := spaceRange{start: sP1, end: eP1}
		pair2 := spaceRange{start: sP2, end: eP2}

		cleaningAreaPairs = append(cleaningAreaPairs, []spaceRange{pair1, pair2})
	}
	readFile.Close()

	return cleaningAreaPairs
}

type spaceRange struct {
	start int
	end   int
}
