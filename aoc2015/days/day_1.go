package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Day1() {
	fmt.Println("Day 1")

	// load exercise data
	santaMap := LoadSantaMapData("aoc2015/inputs/day_1.log")

	// where does santa ends up in ?
	fmt.Println("Q1. To what floor do the instructions take Santa?")
	fmt.Printf("A1. To floor: %d\n", santaMap.floorsUp-santaMap.floorsDown)

	// which position makes santa enter the basement
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

	fmt.Println("Q2. What is the position of the character that causes Santa to first enter the basement?")
	fmt.Printf("A2. Position: %d\n", basementIndex)
}

type santaMap struct {
	journey    []string
	floorsUp   int
	floorsDown int
}

func LoadSantaMapData(fileName string) *santaMap {
	readFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
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
