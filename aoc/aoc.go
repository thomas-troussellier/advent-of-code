package aoc

import "fmt"

type Event interface {
	Run(day int)
}

type AoCEvent struct {
	Day_func_map map[int]func()
	EventYear    string
	Dir          string
}

func (aoc AoCEvent) Run(day int) {
	fmt.Println("~~~~~~~~~~~~~~~~~~~")
	fmt.Printf("Advent of code %s\n", aoc.EventYear)
	fmt.Println("~~~~~~~~~~~~~~~~~~~")

	if day == 0 {
		for _, f := range aoc.Day_func_map {
			f()
		}
	} else {
		aoc.Day_func_map[day]()
	}
}
