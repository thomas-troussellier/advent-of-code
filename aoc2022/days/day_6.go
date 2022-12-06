package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Day6() {
	fmt.Println("Day 6")

	fmt.Println("Q1. How many characters need to be processed before the first start-of-packet marker is detected?")
	signalStart := loadSignalData("aoc2022/inputs/day_6.log", 4)
	fmt.Printf("A1. Characters needing to be processed before the first start-of-packet marker: %d\n", signalStart)

	fmt.Println("Q2. How many characters need to be processed before the first start-of-message marker is detected?")
	signalStart = loadSignalData("aoc2022/inputs/day_6.log", 14)
	fmt.Printf("A2. Characters needing to be processed before the first start-of-message is detected: %d\n", signalStart)
}

func loadSignalData(fileName string, diffChars int) int {
	readFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
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
			return signalStart
		}

		tempSignal = tempSignal[firstIndex+1:]
	}

	readFile.Close()

	return signalStart
}
