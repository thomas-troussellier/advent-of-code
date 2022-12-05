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

func Day1() {
	fmt.Println("Day 1")

	// load exercise data
	elves := LoadElvesData("aoc2022/inputs/day_1.log")

	// sort list by highest load first
	elves.Sort(func(i, j int) bool {
		return elves[i].total > elves[j].total
	})

	// which elf is carrying the most calories ?
	fmt.Println("Q1. Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?")
	fmt.Printf("A1. Elf with the most calories: n° %s, with %d total calories\n", elves[0].name, elves[0].total)

	// which 3 elves are carrying the most calories ? How much is that ?
	cumulativeLoad := 0
	var chadElves []string
	for _, e := range elves[:3] {
		cumulativeLoad += e.total
		chadElves = append(chadElves, e.name)
	}

	fmt.Println("Q1. Find the top three Elves carrying the most Calories. How many Calories are those Elves carrying in total?")
	fmt.Printf("A2. Top three elves are %s, with %d total calories\n", strings.Join(chadElves, ","), cumulativeLoad)
}

type elf struct {
	name  string
	load  []int
	total int
}

func (e elf) String() string {
	return fmt.Sprintf("name: %s, total: %d, load: %v\n", e.name, e.total, e.load)
}

func (e *elf) AddLoad(calories int) {
	e.load = append(e.load, calories)
}

type elves []*elf

func (e elves) Sort(sortFunc func(i, j int) bool) {
	sort.SliceStable(e, sortFunc)
}

func NewElf(name string) *elf {
	return &elf{name: name}
}

func LoadElvesData(fileName string) elves {
	readFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	elves := make(elves, 0, 10)
	elvesCount := 0
	currentElf := NewElf("1")

	for fileScanner.Scan() {
		currentLine := fileScanner.Text()

		if len(currentLine) == 0 {
			elvesCount += 1
			elves = append(elves, currentElf)
			currentElf = NewElf(strconv.Itoa(elvesCount + 1))
		} else {
			intLoad, err := strconv.Atoi(currentLine)
			if err != nil {
				log.Fatal("invalid input")
			}
			currentElf.AddLoad(intLoad)
			currentElf.total += intLoad
		}
	}
	readFile.Close()

	return elves
}
