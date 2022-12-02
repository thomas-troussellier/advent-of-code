package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

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
		fmt.Println(err)
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

func main() {
	fmt.Println("~~~~~~~~~~~~~~~~~~~")
	fmt.Println("Advent of code 2022")
	fmt.Println("~~~~~~~~~~~~~~~~~~~")

	Day1()
}
