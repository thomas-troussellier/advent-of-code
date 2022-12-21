package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day13() {

	fmt.Println("Day 13")
	pairs := loadPacketsPairs("aoc2022/inputs/day_13.log")

	nbSorted := checkPackets(pairs)

	fmt.Println("Q1. What is the sum of the indices of those pairs?")
	fmt.Printf("A1. the sum of the indices of sorted pairs: %d\n", nbSorted)

	decoderKey := sortPackets(pairs)

	fmt.Println("Q2. What is the decoder key for the distress signal?")
	fmt.Printf("A2. decoder key: %d\n", decoderKey)

}

func checkPackets(pairs []*pairs) (total int) {
	total = 0
	for i := 0; i < len(pairs); i++ {
		p := pairs[i]
		compa := compare(p.left, p.right)
		if compa < 0 {
			total += (1 + i)
		}
	}

	return
}

func sortPackets(pairs []*pairs) (total int) {
	total = 0
	pz := make([]any, 0)
	for i := 0; i < len(pairs); i++ {
		p := pairs[i]
		pz = append(pz, p.left, p.right)
	}

	divPacket1 := []any{[]any{2}}
	divPacket2 := []any{[]any{6}}

	pz = append(pz, divPacket1)
	pz = append(pz, divPacket2)

	sort.Slice(pz, func(i, j int) bool {
		return compare(pz[i].([]any), pz[j].([]any)) < 0
	})

	decoder := 1

	for i := range pz {
		if fmt.Sprint(pz[i]) == "[[2]]" || fmt.Sprint(pz[i]) == "[[6]]" {
			decoder *= (i + 1)
		}
	}

	return decoder
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
			currentPair.left = parseInput(currentLine).([]any)
		} else if currentPair.right == nil {
			currentPair.right = parseInput(currentLine).([]any)
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
	left  []any
	right []any
}

func (p pairs) String() string {
	return fmt.Sprintf("l:%s r:%s\n", p.left, p.right)
}

func parseInput(input string) any {
	index := 0
	finalValue := make([]any, 0)
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
	op := 1
	ind := -1
	for i, r := range input[1:] {
		if op == 0 {
			break
		}
		if r == '[' {
			op++
			continue
		}
		if r == ']' {
			op--
			ind = i
			continue
		}
	}
	return ind + 2
}

func compare(l1, l2 []any) int {
	for index := 0; index < len(l1); index++ {
		if index >= len(l2) {
			return 1
		}

		var comp int

		switch v1 := l1[index].(type) {
		case int:
			// check against l2 value
			switch v2 := l2[index].(type) {
			case int:
				comp = v1 - v2
			case []any:
				comp = compare([]any{v1}, v2)
			}
		case []any:
			// check against l2 value
			switch v2 := l2[index].(type) {
			case int:
				comp = compare(v1, []any{v2})
			case []any:
				comp = compare(v1, v2)
			}
		}

		if comp < 0 {
			return -1
		}
		if comp > 0 {
			return 1
		}
		continue
	}

	if len(l1) < len(l2) {
		return -1
	}

	return 0
}
