package aoc2015

import (
	"advent-of-code/aoc"
)

var Aoc2015 = aoc.AoCEvent{
	Day_func_map: map[int]func(){
		1: day1,
		2: day2,
		3: day3,
		4: day4,
	},
	EventYear: "2015",
	Dir:       "aoc2015",
}
