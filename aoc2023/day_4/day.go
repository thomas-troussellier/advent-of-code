package day_4

import (
	"advent-of-code/aoc"
	"advent-of-code/aoc/utils"
	"bufio"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
)

type day struct {
	inputFile string
}

var _ aoc.EventRunner = (*day)(nil)

func Create() aoc.EventRunner {
	return create("aoc2023/day_4/input.txt")
}

func create(inputFile string) *day {
	return &day{
		inputFile: inputFile,
	}
}

func (d *day) Execute() {
	log.Println("Day 4")
	log.Println("Q1:", d.Question1())
	log.Println("Q2:", d.Question2())
}

func (d *day) Question1() string {
	n, o := d.loadData()
	sum := 0
	for i, nbs := range o {
		res := make([]int, 0)
		for _, nb := range nbs {
			if slices.Contains(n[i], nb) {
				res = append(res, nb)
			}
		}
		if len(res) > 0 {
			sum += int(math.Pow(2, float64(len(res)-1)))
		}
	}
	return strconv.Itoa(sum)
}

func (d *day) Question2() string {
	cards, ownedNumbers := d.loadData()
	resMap := make(map[int]int)
	oOrd := make([]int, 0)
	for v := range ownedNumbers {
		oOrd = append(oOrd, v)
	}
	slices.Sort(oOrd)
	for _, ind := range oOrd {
		matchingSymbols := make([]int, 0)
		for _, nb := range ownedNumbers[ind] {
			if slices.Contains(cards[ind], nb) {
				matchingSymbols = append(matchingSymbols, nb)
			}
		}

		resMap[ind] += 1
		for j := range matchingSymbols {
			if _, ok := cards[ind+j+1]; ok {
				resMap[ind+j+1] += (1 * resMap[ind])
			}

		}
	}
	sum := 0
	for _, v := range resMap {
		sum += v
	}
	return strconv.Itoa(sum)
}

func (d *day) loadData() (map[int][]int, map[int][]int) {
	file := utils.LoadInput(d.inputFile)

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	nb := make(map[int][]int, 0)
	own := make(map[int][]int, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		line = strings.Join(strings.Fields(line), " ")
		base := strings.Split(line, ":")
		ind, _ := strconv.Atoi(strings.Split(base[0], " ")[1])
		for i, l := range strings.Split(base[1], "|") {
			for _, s := range strings.Split(l, " ") {
				if s == "" {
					continue
				}
				n, _ := strconv.Atoi(s)
				if i == 0 {
					nb[ind] = append(nb[ind], n)
				} else {
					own[ind] = append(own[ind], n)
				}
			}
		}
	}

	return nb, own
}
