package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Day3() {
	fmt.Println("Day 3")

	// load exercise data
	rumsacks := LoadRumsackData("aoc2022/inputs/day_3.log")

	fmt.Println("Q1. What is the sum of the priorities of those item types?")
	totalPriorities := findPriorities1(rumsacks)

	fmt.Printf("A1. Total score will be: %d\n", totalPriorities)

	fmt.Println("Q2. What is the sum of the priorities of those item types?")
	totalPriorities = findPriorities2(rumsacks)

	fmt.Printf("A2. Total score will be: %d\n", totalPriorities)
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

func LoadRumsackData(fileName string) []string {
	readFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
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
