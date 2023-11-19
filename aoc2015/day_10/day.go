package day_10

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
	return create("aoc2015/day_10/input.txt")
}

func create(inputFile string) *day {
	return &day{
		inputFile: inputFile,
	}
}

func (d *day) Execute() {
	log.Println("Day 10")
	log.Println("Q1:", d.Question1())
	log.Println("Q2:", d.Question2())
}

func (d *day) Question1() string {
	lookSay := loadLookSay(d.inputFile)

	return strconv.Itoa(len(computeLookSay(lookSay, 40)))
}

func (d *day) Question2() string {
	lookSay := loadLookSay(d.inputFile)

	return strconv.Itoa(len(computeLookSay(lookSay, 50)))
}

func computeLookSay(input string, repeat int) string {

	tempString := input
	res := make([]string, 0)

	for i := 0; i < repeat; i++ {
		for len(tempString) > 0 {
			newInput := strings.TrimLeft(tempString, string(tempString[0]))
			occ := len(tempString) - len(newInput)
			res = append(res, strings.Join([]string{strconv.Itoa(occ), string(tempString[0])}, ""))
			tempString = newInput
		}
		tempString = strings.Join(res, "")
		res = make([]string, 0)
	}

	return tempString
}

func loadLookSay(fileName string) string {
	readFile := utils.LoadInput(fileName)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var lookSayString string
	for fileScanner.Scan() {
		lookSayString = fileScanner.Text()
	}

	readFile.Close()

	return lookSayString
}
