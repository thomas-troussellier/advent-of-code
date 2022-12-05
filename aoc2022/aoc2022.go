package aoc2022

import "advent-of-code/aoc"

var Aoc2022 = aoc.AoCEvent{
	Day_func_map: map[int]func(){
		1: day1,
		2: day2,
		3: day3,
		4: day4,
		5: day5,
	},
	EventYear: "2022",
	Dir:       "aoc2022",
}
