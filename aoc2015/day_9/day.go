package day_9

import (
	"advent-of-code/aoc"
	"advent-of-code/aoc/utils"
	"bufio"
	"fmt"
	"log"
	"strconv"
)

type day struct {
	inputFile string
}

var _ aoc.EventRunner = (*day)(nil)

func Create() aoc.EventRunner {
	return create("aoc2015/day_9/input.txt")
}

func create(inputFile string) *day {
	return &day{
		inputFile: inputFile,
	}
}

func (d *day) Execute() {
	log.Println("Day 9")
	log.Println("Q1:", d.Question1())
	log.Println("Q2:", d.Question2())
}

func (d *day) Question1() string {
	distances := loadDistances(d.inputFile)

	return strconv.Itoa(computeDist(distances, true))
}

func (d *day) Question2() string {
	distances := loadDistances(d.inputFile)

	return strconv.Itoa(computeDist(distances, false))
}

func computeDist(distances map[string]map[string]int, short bool) int {
	dist := 0
	dests := make([]string, 0)

	for from := range distances {
		// for each town, try to find the most visited places + shortest dist

		tempVisited := make([]string, 0)
		tempDist := 0

		tempVisited, tempDist = getDistFrom(distances, from, tempVisited, tempDist, short)

		//fmt.Printf("From: %s, Shortest travel: %d, road: %v\n", from, tempDist, tempVisited)
		if short && len(tempVisited) >= len(dests) && ((dist == 0) || (tempDist <= dist)) {
			dist = tempDist
			dests = tempVisited
		} else if !short && len(tempVisited) >= len(dests) && ((dist == 0) || (tempDist >= dist)) {
			dist = tempDist
			dests = tempVisited
		}
	}
	return dist
}

func getDistFrom(distances map[string]map[string]int, from string, visited []string, distance int, short bool) ([]string, int) {
	visited = append(visited, from)
	if to, dist := getDist(distances[from], visited, short); to != "" {
		distance += dist
		return getDistFrom(distances, to, visited, distance, short)
	}

	return visited, distance
}

func getDist(from map[string]int, visited []string, short bool) (to string, dist int) {
	for t, d := range from {
		if contains(visited, t) {
			continue
		}

		if short && (dist == 0 || d < dist) {
			dist = d
			to = t
		} else if !short && (dist == 0 || d > dist) {
			dist = d
			to = t
		}
	}

	return to, dist
}

func contains(list []string, s string) bool {
	for _, st := range list {
		if st == s {
			return true
		}
	}
	return false
}

func loadDistances(fileName string) map[string]map[string]int {
	readFile := utils.LoadInput(fileName)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	distances := make(map[string]map[string]int)

	for fileScanner.Scan() {
		direction := fileScanner.Text()
		var from, to string
		var dist int
		fmt.Sscanf(direction, "%s to %s = %d", &from, &to, &dist)

		if _, ok := distances[from]; !ok {
			distances[from] = make(map[string]int)
		}
		distances[from][to] = dist
		if _, ok := distances[to]; !ok {
			distances[to] = make(map[string]int)
		}
		distances[to][from] = dist
	}

	readFile.Close()

	return distances
}
