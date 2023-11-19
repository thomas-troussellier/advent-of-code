package day_11

import (
	"advent-of-code/aoc"
	"advent-of-code/aoc/utils"
	"bufio"
	"log"
	"strings"
)

type day struct {
	inputFile string
}

var _ aoc.EventRunner = (*day)(nil)

func Create() aoc.EventRunner {
	return create("aoc2015/day_11/input.txt")
}

func create(inputFile string) *day {
	return &day{
		inputFile: inputFile,
	}
}

func (d *day) Execute() {
	log.Println("Day 11")
	log.Println("Q1:", d.Question1())
	log.Println("Q2:", d.Question2())
}

func (d *day) Question1() string {
	password := loadPassword(d.inputFile)

	return computeNextPassword(password, false)
}

func (d *day) Question2() string {
	password := loadPassword(d.inputFile)
	password = computeNextPassword(password, true)
	return computeNextPassword(password, true)
}

func computeNextPassword(password string, forceRegen bool) string {
	valid, allowed, fIndex := checkPassRules(password)

	if valid && !forceRegen {
		return password
	}

	nextPass := password

	// make replace first forbidden char from pass
	if !allowed {
		nextPass = strings.Replace(nextPass, string(nextPass[fIndex]), string(nextPass[fIndex]+1), 1)
		nextPass = nextPass[:fIndex+1] + strings.Repeat("a", len(nextPass[fIndex+1:]))
		return computeNextPassword(nextPass, false)
	}

	nextPass = getNextIncPass(nextPass)
	return computeNextPassword(nextPass, false)
}

func getNextIncPass(input string) string {
	endChar := input[len(input)-1]

	if endChar == 'z' {
		return getNextIncPass(input[:len(input)-1]) + "a"
	}

	return input[:len(input)-1] + string(endChar+1)
}

func checkPassRules(pass string) (isValid, allowed bool, fIndex int) {
	var length, seq bool
	var dupCount int

	if len(pass) == 8 {
		length = true
	}

	if strings.ContainsAny(pass, "iol") {
		fIndex = strings.IndexAny(pass, "iol")
	} else {
		allowed = true
	}

	for i := 0; i < len(pass)-2; i++ {
		if pass[i] == pass[i+1]-1 && pass[i+1] == pass[i+2]-1 {
			seq = true
			break
		}
	}

	for i := 0; i < len(pass)-1; i++ {
		if pass[i] == pass[i+1] {
			dupCount++
			i += 1
		}
		if dupCount == 2 {
			break
		}
	}

	isValid = length && allowed && seq && (dupCount == 2)

	return
}

func loadPassword(fileName string) string {
	readFile := utils.LoadInput(fileName)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var password string
	for fileScanner.Scan() {
		password = fileScanner.Text()
	}

	readFile.Close()

	return password
}
