package day_1

import (
	"advent-of-code/aoc"
	"advent-of-code/aoc/utils"
	"bufio"
	"fmt"
	"log"
	"sort"
	"strconv"
)

type day struct {
	inputFile string
}

var _ aoc.EventRunner = (*day)(nil)

func Create() aoc.EventRunner {
	return create("aoc2022/day_1/input.txt")
}

func create(inputFile string) *day {
	return &day{
		inputFile: inputFile,
	}
}

func (d *day) Execute() {
	log.Println("Day 1")
	log.Println("Q1:", d.Question1())
	log.Println("Q2:", d.Question2())
}

func (d *day) Question1() string {

	elves := loadElvesData(d.inputFile)

	// sort list by highest load first
	elves.Sort(func(i, j int) bool {
		return elves[i].total > elves[j].total
	})

	return strconv.Itoa(elves[0].total)
}

func (d *day) Question2() string {

	elves := loadElvesData(d.inputFile)

	// sort list by highest load first
	elves.Sort(func(i, j int) bool {
		return elves[i].total > elves[j].total
	})

	cumulativeLoad := 0
	for _, e := range elves[:3] {
		e := e
		cumulativeLoad += e.total
	}

	return strconv.Itoa(cumulativeLoad)
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

func newElf(name string) *elf {
	return &elf{name: name}
}

func loadElvesData(fileName string) elves {
	readFile := utils.LoadInput(fileName)
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	elves := make(elves, 0)
	elves = append(elves, newElf("1"))

	for fileScanner.Scan() {
		currentLine := fileScanner.Text()
		index := len(elves) - 1
		if len(currentLine) == 0 {
			elves = append(elves, newElf(strconv.Itoa(len(elves)+1)))
		} else {
			intLoad, err := strconv.Atoi(currentLine)
			if err != nil {
				log.Fatal("invalid input")
			}
			elves[index].AddLoad(intLoad)
			elves[index].total += intLoad
		}
	}

	return elves
}
