package aoc2015

import (
	"advent-of-code/aoc"
	"advent-of-code/aoc2015/days"
)

var Aoc2015 = aoc.AoCEvent{
	Day_func_map: map[int]func(){
		1:  days.Day1,
		2:  days.Day2,
		3:  days.Day3,
		4:  days.Day4,
		5:  days.Day5,
		6:  days.Day6,
		7:  days.Day7,
		8:  days.Day8,
		9:  days.Day9,
		10: days.Day10,
		11: days.Day11,
	},
	EventYear: "2015",
	Dir:       "aoc2015",
}
