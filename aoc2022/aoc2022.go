package aoc2022

import "advent-of-code/aoc"

var Aoc2022 = aoc.AoCEvent{
	Day_func_map: map[int]func(){
		1: day1,
		2: day2,
	},
	EventYear: "2022",
	Dir:       "aoc2022",
}
