package day_13

import (
	"advent-of-code/aoc"
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type day struct {
	inputFile string
}

var _ aoc.EventRunner = (*day)(nil)

func Create() aoc.EventRunner {
	return create("aoc2022/day_13/input.txt")
}

func create(inputFile string) *day {
	return &day{
		inputFile: inputFile,
	}
}

func (d *day) Execute() {
	log.Println("Day 13")
	log.Println("Q1:", d.Question1())
	log.Println("Q2:", d.Question2())
}

func (d *day) Question1() string {
	pairs := loadPacketsPairs(d.inputFile)

	nbSorted := checkPackets(pairs)

	return strconv.Itoa(nbSorted)
}

func (d *day) Question2() string {
	pairs := loadPacketsPairs(d.inputFile)

	decoderKey := sortPackets(pairs)

	return strconv.Itoa(decoderKey)
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
