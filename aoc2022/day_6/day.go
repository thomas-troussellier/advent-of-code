package day_6

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
	return create("aoc2022/day_6/input.txt")
}

func create(inputFile string) *day {
	return &day{
		inputFile: inputFile,
	}
}

func (d *day) Execute() {
	log.Println("Day 6")
	log.Println("Q1:", d.Question1())
	log.Println("Q2:", d.Question2())
}

func (d *day) Question1() string {

	return loadSignalData(d.inputFile, 4)
}

func (d *day) Question2() string {

	return loadSignalData(d.inputFile, 14)
}

func loadSignalData(fileName string, diffChars int) string {
	readFile := utils.LoadInput(fileName)

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanRunes)

	signalStart := 0
	tempSignal := make([]string, 0)

	for fileScanner.Scan() {
		currentChar := fileScanner.Text()
		signalStart += 1
		tempSignal = append(tempSignal, currentChar)

		if len(tempSignal) != diffChars {
			continue
		}

		trueSignal := true
		firstIndex := 0
		joined := strings.Join(tempSignal, "")
		for _, s := range tempSignal {
			if strings.Count(joined, s) > 1 {
				trueSignal = false
				firstIndex = strings.Index(joined, s)
				break
			}
		}

		if trueSignal {
			return strconv.Itoa(signalStart)
		}

		tempSignal = tempSignal[firstIndex+1:]
	}

	readFile.Close()

	return strconv.Itoa(signalStart)
}
