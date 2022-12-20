package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day13() {

	fmt.Println("Day 13")
	pairs := loadPacketsPairs("aoc2022/inputs/day_13.log")

	nbSorted := checkPackets(pairs)

	fmt.Println("Q1. What is the sum of the indices of those pairs?")
	fmt.Printf("A1. the sum of the indices of sorted pairs: %d\n", nbSorted)

}

func checkPackets(pairs []*pairs) (total int) {
	for i, p := range pairs {
		if p.compare() {
			total = total + (i + 1)
		}
	}

	return
}

func loadPacketsPairs(fileName string) []*pairs {
	readFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	pairList := make([]*pairs, 0)

	currentPair := pairs{}

	for fileScanner.Scan() {
		currentLine := fileScanner.Text()
		if len(currentLine) == 0 {
			currentLine = ""
		}

		if currentPair.left == nil {
			currentPair.left = parseInput(currentLine).([]interface{})
		} else if currentPair.right == nil {
			currentPair.right = parseInput(currentLine).([]interface{})
		} else {
			temp := currentPair
			pairList = append(pairList, &temp)
			currentPair = pairs{}
		}
	}
	pairList = append(pairList, &currentPair)

	readFile.Close()

	return pairList
}

type pairs struct {
	left  []interface{}
	right []interface{}
}

func (p pairs) String() string {
	return fmt.Sprintf("l:%s r:%s\n", p.left, p.right)
}

func (p pairs) compare() bool {
	return compareList(p.left, p.right)
}

func compareList(pLeft []interface{}, pRight []interface{}) bool {

	if len(pLeft) > len(pRight) {
		return !compareList(pRight, pLeft)
	}

	index := 0
	for index < len(pLeft) {
		rightOrder := true
		switch v := pLeft[index].(type) {
		case []interface{}:
			switch v2 := pRight[index].(type) {
			case []interface{}:
				rightOrder = compareList(v, v2)
			case int:
				rightOrder = compareList(v, []interface{}{v2})
			}
		case int:
			switch v2 := pRight[index].(type) {
			case []interface{}:
				rightOrder = compareList([]interface{}{v}, v2)
			case int:
				rightOrder = compareInt(v, v2)
			}
		}

		if !rightOrder {
			return false
		}
		index++
	}

	return true
}

func compareInt(v, v2 int) bool {
	return v <= v2
}

func parseInput(input string) interface{} {
	index := 0
	finalValue := make([]interface{}, 0)
	for index < len(input) {
		switch input[index] {
		case '[':
			end := readList(input[index:])
			finalValue = append(finalValue, parseInput(input[index+1:index+end]))
			index += end
		case ']':
			index++
		case ',':
			index++
		default:
			nextSep := strings.IndexAny(input[index:], ",[]")
			if nextSep < 0 {
				index++
				nextSep = len(input[index:])
			} else if nextSep == 0 {
				nextSep++
			}
			nb, _ := strconv.Atoi(input[index : index+nextSep])
			finalValue = append(finalValue, nb)
			index += nextSep
		}
	}

	return finalValue
}

func readList(input string) int {
	nextClosing := strings.Index(input, "]")
	nextOpening := strings.Index(input[1:], "[")

	if (nextOpening < 0) || (nextClosing < nextOpening) {
		return nextClosing + 1
	}

	countOpening := strings.Count(input, "[")

	tIndex := 0
	for countOpening > 0 {
		if tIndex == len(input) {
			return tIndex
		}
		if input[tIndex] == ']' {
			countOpening--
		}
		tIndex++
	}

	return tIndex
}
