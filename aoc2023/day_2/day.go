package day_2

import (
	"advent-of-code/aoc"
	"advent-of-code/aoc/utils"
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type day struct {
	inputFile string
}

var _ aoc.EventRunner = (*day)(nil)

func Create() aoc.EventRunner {
	return create("aoc2023/day_2/input.txt")
}

func create(inputFile string) *day {
	return &day{
		inputFile: inputFile,
	}
}

func (d *day) Execute() {
	log.Println("Day 2")
	log.Println("Q1:", d.Question1())
	log.Println("Q2:", d.Question2())
}

func (d *day) Question1() string {
	games := d.loadData()
	return strconv.Itoa(d.sumIds(games, 12, 13, 14))
}

func (d *day) Question2() string {
	games := d.loadData()

	return strconv.Itoa(d.computeTotalPower(games))
}

func (d *day) computeTotalPower(games []game) int {
	sumPowers := 0
	for _, g := range games {
		sumPowers += g.power()
	}

	return sumPowers
}

func (d *day) sumIds(games []game, r, g, b int) int {
	sum := 0
	for _, ga := range games {
		if d.possibleInput(ga, r, g, b) {
			sum += ga.id
		}
	}
	return sum
}

func (d *day) possibleInput(ga game, r, g, b int) bool {
	return ((ga.getColor("red") <= r) && (ga.getColor("green") <= g) && (ga.getColor("blue") <= b))
}

func (g game) getColor(color string) int {
	if v, ok := g.colors[color]; ok {
		return v
	}
	return 0
}

type game struct {
	colors map[string]int
	id     int
}

func (g game) power() int {
	if len(g.colors) == 0 {
		return 0
	}
	pow := 1
	for _, v := range g.colors {
		pow *= v
	}
	return pow
}

func (d *day) loadData() []game {
	file := utils.LoadInput(d.inputFile)

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	games := make([]game, 0)
	for fileScanner.Scan() {
		in := strings.Split(fileScanner.Text(), ":")
		g := game{
			colors: make(map[string]int),
		}
		fmt.Sscanf(in[0], "Game %d", &g.id)
		for _, s := range strings.Split(in[1], ";") {
			for _, r := range strings.Split(s, ",") {
				vl := strings.Split(strings.TrimSpace(r), " ")
				colorValue, _ := strconv.Atoi(vl[0])
				colorName := vl[1]
				currentValue := g.colors[colorName]
				if currentValue < colorValue {
					g.colors[colorName] = colorValue
				}
			}
		}
		games = append(games, g)
	}

	return games
}
