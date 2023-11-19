package aoc

import "log"

type Event interface {
	Run(day int)
}

type EventRunner interface {
	Execute()
	Question1() string
	Question2() string
}

type AoCEvent struct {
	Day_func_map map[int]EventRunner
	EventYear    string
}

func (aoc AoCEvent) Run(day int) {
	log.Println("~~~~~~~~~~~~~~~~~~~")
	log.Printf("Advent of code %s\n", aoc.EventYear)
	log.Println("~~~~~~~~~~~~~~~~~~~")

	if day == 0 {
		for _, day := range aoc.Day_func_map {
			day.Execute()
		}
	} else {
		if dayFunc, ok := aoc.Day_func_map[day]; ok {
			dayFunc.Execute()
		} else {
			log.Fatalf("execution unavailable for day %d", day)
		}
	}
}
