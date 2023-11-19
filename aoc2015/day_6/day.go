package day_6

import (
	"advent-of-code/aoc"
	"advent-of-code/aoc/utils"
	"bufio"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

type day struct {
	inputFile string
}

var _ aoc.EventRunner = (*day)(nil)

func Create() aoc.EventRunner {
	return create("aoc2015/day_6/input.txt")
}

func create(inputFile string) *day {
	return &day{
		inputFile: inputFile,
	}
}

func (d *day) Execute() {
	log.Println("Day 6")
	log.Println("Q1:", d.Question1())
	log.Println("Q2:", d.Question2())
}

func (d *day) Question1() string {
	instructions := loadLightInstr(d.inputFile)

	var grid [1000][1000]bool

	for _, instruction := range instructions {
		lx, rx, ly, ry, instr := convertInstruction(instruction)

		for i := lx; i < rx+1; i++ {
			for j := ly; j < ry+1; j++ {
				switch instr {
				case on:
					grid[i][j] = true
				case off:
					grid[i][j] = false
				case toggle:
					grid[i][j] = !grid[i][j]
				}

			}
		}
	}
	lit := 0

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if grid[i][j] {
				lit++
			}
		}
	}

	return strconv.Itoa(lit)
}

func (d *day) Question2() string {
	instructions := loadLightInstr(d.inputFile)

	var grid2 [1000][1000]int
	for _, instruction := range instructions {
		lx, rx, ly, ry, instr := convertInstruction(instruction)

		for i := lx; i < rx+1; i++ {
			for j := ly; j < ry+1; j++ {
				switch instr {
				case on:
					grid2[i][j] += 1
				case off:
					if grid2[i][j]-1 >= 0 {
						grid2[i][j] -= 1
					}
				case toggle:
					grid2[i][j] += 2
				}

			}
		}
	}

	brightness := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			brightness += grid2[i][j]
		}
	}

	return strconv.Itoa(brightness)
}

type order uint8

const (
	on order = iota
	off
	toggle
)

func convertInstruction(instruction string) (lx, rx, ly, ry int, instr order) {
	var next string
	if strings.HasPrefix(instruction, "turn on") {
		instr = on
		next = instruction[8:]
	} else if strings.HasPrefix(instruction, "turn off") {
		instr = off
		next = instruction[9:]
	} else {
		instr = toggle
		next = instruction[7:]
	}

	var tmpLx, tmpLy, tmpRx, tmpRy int
	_, err := fmt.Sscanf(next, "%d,%d through %d,%d", &tmpLx, &tmpLy, &tmpRx, &tmpRy)
	if err != nil {
		fmt.Println(err)
	}

	x := []int{tmpLx, tmpRx}
	sort.Ints(x)
	y := []int{tmpLy, tmpRy}
	sort.Ints(y)

	return x[0], x[1], y[0], y[1], instr
}

func loadLightInstr(fileName string) []string {
	readFile := utils.LoadInput(fileName)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	list := make([]string, 0)

	for fileScanner.Scan() {
		instruction := fileScanner.Text()

		list = append(list, instruction)

	}

	readFile.Close()

	return list
}
