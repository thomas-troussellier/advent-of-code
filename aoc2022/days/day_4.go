package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Day4() {
	fmt.Println("Day 4")

	// load exercise data
	pairsData := LoadSpacePairsData("aoc2022/inputs/day_4.log")

	fmt.Println("Q1. In how many assignment pairs does one range fully contain the other?")
	totalContainedPairs := findContained(pairsData)
	fmt.Printf("A1. Fully contained pairs: %d\n", totalContainedPairs)

	fmt.Println("Q2. In how many assignment pairs do the ranges overlap?")
	totalOverlapPairs := findOverlapped(pairsData)
	fmt.Printf("A2. Overlapped ranges pairs: %d\n", totalOverlapPairs)
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

func LoadSpacePairsData(fileName string) [][]spaceRange {
	readFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
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
