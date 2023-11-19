package main

import (
	"advent-of-code/aoc"
	"advent-of-code/aoc/utils"
	"advent-of-code/aoc2015"
	"advent-of-code/aoc2022"
	"advent-of-code/aoc2023"
	"log"
	"os"
	"strconv"
)

var event_func_map = map[string]aoc.Event{
	"2015": aoc2015.Aoc2015,
	"2022": aoc2022.Aoc2022,
	"2023": aoc2023.Aoc2023,
}

func main() {
	inputs := os.Args

	if len(inputs) < 2 {
		log.Fatal("missing args. usage: go run main.go <execute|create> [year] [day]")
	}

	switch inputs[1] {

	case "execute":
		if len(inputs) == 2 {
			runAll()
		} else {
			day := "0"
			if len(inputs) >= 4 {
				day = inputs[3]
			}
			runYear(inputs[2], day)
		}
	case "create":
		if len(inputs) != 4 {
			log.Fatal("usage: go run main.go create <year> <day>")
		}
		create(inputs[2], inputs[3])
	default:
		log.Fatal("expected commands are: 'execute' or 'create'")
	}
}

func runAll() {
	log.Println("No args given, run all advent of code events")
	for _, event := range event_func_map {
		event.Run(0)
	}
}

func runYear(year string, day string) {
	dayToRun, _ := strconv.Atoi(day)
	event_func_map[year].Run(dayToRun)
}

func create(year, day string) {
	y, _ := strconv.Atoi(year)
	d, _ := strconv.Atoi(day)
	utils.GetInputFromSite(y, d)
	utils.CopyTemplates(y, d)
}
