package main

import (
	"advent-of-code/aoc"
	"advent-of-code/aoc2015"
	"advent-of-code/aoc2022"
	"fmt"
	"os"
	"strconv"
)

var event_func_map = map[string]aoc.Event{
	"2015": aoc2015.Aoc2015,
	"2022": aoc2022.Aoc2022,
}

func main() {
	inputs := os.Args

	var dayToRun int = 0

	if len(inputs) == 1 {
		fmt.Println("No args given, run all advent of code events")
		for _, event := range event_func_map {
			event.Run(dayToRun)
		}
	} else if len(inputs) >= 2 {
		if len(inputs) == 3 {
			dayToRun, _ = strconv.Atoi(inputs[2])
		}
		event_func_map[inputs[1]].Run(dayToRun)
	}
}
