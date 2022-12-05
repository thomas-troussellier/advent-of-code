package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func Day8() {
	fmt.Println("Day 8")

	// load exercise data
	matchsticks := loadMatchsticks("aoc2015/inputs/day_8.log")

	fmt.Println("Q1. What is the number of characters of code for string literals minus the number of characters in memory for the values of the strings in total for the entire file?")
	total := parseData(matchsticks)
	fmt.Printf("A1. sol: %d\n", total)

	fmt.Println("Q2. Find the total number of characters to represent the newly encoded strings minus the number of characters of code in each original string literal")
	total = parseData2(matchsticks)
	fmt.Printf("A2. sol: %d\n", total)
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
	readFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
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
