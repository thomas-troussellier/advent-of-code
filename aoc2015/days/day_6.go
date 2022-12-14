package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func Day6() {
	fmt.Println("Day 6")

	// load exercise data
	instructions := loadLightInstr("aoc2015/inputs/day_6.log")

	var grid [1000][1000]bool

	fmt.Println("Q1. After following the instructions, how many lights are lit?")
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
	fmt.Printf("A1. lights lit: %d\n", lit)

	fmt.Printf("Q2. What is the total brightness of all lights combined after following Santa's instructions?\n")
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
	fmt.Printf("A2. total light brightness: %d\n", brightness)
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
	readFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
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
