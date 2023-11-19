package day_1

import (
	"advent-of-code/aoc"
	"advent-of-code/aoc/utils"
	"bufio"
	"log"
	"strconv"
)

type day struct {
	inputFile string
}

var _ aoc.EventRunner = (*day)(nil)

func Create() aoc.EventRunner {
	return create("aoc2015/day_1/input.txt")
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
	santaMap := loadSantaMapData(d.inputFile)

	return strconv.Itoa(santaMap.floorsUp - santaMap.floorsDown)
}

func (d *day) Question2() string {
	santaMap := loadSantaMapData(d.inputFile)
	basementIndex, currentFloor := 0, 0
	for i := 0; i < len(santaMap.journey); i++ {
		if santaMap.journey[i] == "(" {
			currentFloor += 1
		} else {
			currentFloor -= 1
		}
		if currentFloor == -1 {
			basementIndex = i + 1
			break
		}
	}
	return strconv.Itoa(basementIndex)
}

type santaMap struct {
	journey    []string
	floorsUp   int
	floorsDown int
}

func loadSantaMapData(fileName string) *santaMap {
	readFile := utils.LoadInput(fileName)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanRunes)

	santa := &santaMap{journey: make([]string, 0)}

	for fileScanner.Scan() {
		currentCommand := fileScanner.Text()

		switch currentCommand {
		case "(":
			santa.floorsUp += 1
		case ")":
			santa.floorsDown += 1
		}

		santa.journey = append(santa.journey, currentCommand)
	}
	readFile.Close()

	return santa
}
