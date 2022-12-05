package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Day9() {
	fmt.Println("Day 9")

	// load exercise data
	distances := loadDistances("aoc2015/inputs/day_9.log")

	fmt.Println("Q1. What is the distance of the shortest route?")
	shortestRouteDist := computeDist(distances, true)
	fmt.Printf("A1. shortest route: %d\n", shortestRouteDist)

	fmt.Println("Q2. What is the distance of the longest route?")
	longestRouteDist := computeDist(distances, false)
	fmt.Printf("A2. longest route: %d\n", longestRouteDist)
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

	fmt.Printf("travel: %d, road: %v\n", dist, dests)
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
	readFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
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
