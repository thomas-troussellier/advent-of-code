package day_5

import (
	"advent-of-code/aoc"
	"advent-of-code/aoc/utils"
	"bufio"
	"log"
	"strconv"
	"strings"
)

type day struct {
	inputFile string
}

var _ aoc.EventRunner = (*day)(nil)

func Create() aoc.EventRunner {
	return create("aoc2015/day_5/input.txt")
}

func create(inputFile string) *day {
	return &day{
		inputFile: inputFile,
	}
}

func (d *day) Execute() {
	log.Println("Day 5")
	log.Println("Q1:", d.Question1())
	log.Println("Q2:", d.Question2())
}

func (d *day) Question1() string {
	entries := loadStringList(d.inputFile)

	noice := 0
	for _, e := range entries {
		// 3 vowels ?
		cmpt := 0
		double := false
		for i, c := range e.base {
			if strings.Contains(must_vowels, string(c)) {
				cmpt++
			}
			// double letters ?
			if i < len(e.base)-1 && string(e.base[i+1]) == string(c) {
				double = true
			}
		}

		// forbidden strings ?
		forbidden_free := true
		for _, s := range forbidden_strings {
			if strings.Count(e.base, s) > 0 {
				forbidden_free = false
				break
			}
		}

		e.nice = (cmpt >= 3) && double && forbidden_free

		if e.nice {
			noice++
		}
	}
	return strconv.Itoa(noice)
}

func (d *day) Question2() string {
	entries := loadStringList(d.inputFile)
	noice := 0
	for _, e := range entries {
		// contains a pair of any two letters that appears at least twice in the string without overlapping
		twoLetters := false
		// contains at least one letter which repeats with exactly one letter between them
		repeats := false

		for i := 0; i < len(e.base)-1; i++ {
			if !twoLetters && strings.Count(e.base, e.base[i:i+2]) > 1 {
				twoLetters = true
			}

			nextIndex := strings.Index(e.base[i+1:], string(e.base[i]))
			if !repeats && (nextIndex == 1) {
				repeats = true
			}

			if twoLetters && repeats {
				noice++
				break
			}
		}
	}
	return strconv.Itoa(noice)
}

type entry struct {
	base string
	nice bool
}

func loadStringList(fileName string) []*entry {
	readFile := utils.LoadInput(fileName)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	list := make([]*entry, 0)

	for fileScanner.Scan() {
		toTest := fileScanner.Text()

		list = append(list, &entry{
			base: toTest,
		})

	}

	readFile.Close()

	return list
}

var must_vowels = "aeiou"
var forbidden_strings = []string{"ab", "cd", "pq", "xy"}
