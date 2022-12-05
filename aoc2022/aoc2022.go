package aoc2022

import (
	"advent-of-code/aoc"
	"advent-of-code/aoc2022/days"
)

var Aoc2022 = aoc.AoCEvent{
	Day_func_map: map[int]func(){
		1: days.Day1,
		2: days.Day2,
		3: days.Day3,
		4: days.Day4,
		5: days.Day5,
	},
	EventYear: "2022",
	Dir:       "aoc2022",
}
