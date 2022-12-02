package main

import (
	"fmt"
	"strings"
)

func Day1() {
	fmt.Println("Day 1")

	// load exercise data
	elves := LoadElvesData("day_1.log")

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
