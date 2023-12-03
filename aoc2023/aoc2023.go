package aoc2023

import (
	"advent-of-code/aoc"
	"advent-of-code/aoc2023/day_1"
	"advent-of-code/aoc2023/day_2"
	"advent-of-code/aoc2023/day_3"
)

var Aoc2023 = aoc.AoCEvent{
	Day_func_map: map[int]aoc.EventRunner{
		1: day_1.Create(),
		2: day_2.Create(),
		3: day_3.Create(),
		//4:  day_4.Create(),
		//5:  day_5.Create(),
		//6:  day_6.Create(),
		//7:  day_7.Create(),
		//8:  day_8.Create(),
		//9:  day_9.Create(),
		//10: day_10.Create(),
		//11: day_11.Create(),
		//12: day_12.Create(),
		//13: day_13.Create(),
		//14: day_14.Create(),
		//15: day_15.Create(),
	},
	EventYear: "2023",
}
