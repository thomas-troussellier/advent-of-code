package aoc2022

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func day5() {
	fmt.Println("Day 5")

	// load exercise data
	crates, moves := LoadCrateData("aoc2022/day_5.log")

	fmt.Println("Q1. After the rearrangement procedure completes, what crate ends up on top of each stack?")
	cratesOnTop := computeMoves9000(crates, moves)
	fmt.Printf("A1. crates on top of each stack: %s\n", cratesOnTop)

	// load exercise data
	crates, moves = LoadCrateData("aoc2022/day_5.log")
	fmt.Println("Q2. After the rearrangement procedure completes, what crate ends up on top of each stack?")
	cratesOnTop = computeMoves9001(crates, moves)
	fmt.Printf("A1. crates on top of each stack: %s\n", cratesOnTop)
}

func computeMoves9000(crates [][]string, moves []instruction) string {
	for _, instr := range moves {
		for i := 0; i < instr.nb; i++ {
			crateFrom := crates[instr.from-1][1:]
			crateTo := append(make([]string, 0), crates[instr.from-1][0])
			crateTo = append(crateTo, crates[instr.to-1]...)

			crates[instr.to-1] = crateTo
			crates[instr.from-1] = crateFrom
		}
	}

	tops := make([]string, 0)
	for _, crate := range crates {
		tops = append(tops, crate[0])
	}

	return strings.Join(tops, "")
}

func computeMoves9001(crates [][]string, moves []instruction) string {
	for _, instr := range moves {
		crateFrom := crates[instr.from-1][instr.nb:]
		crateTo := append(make([]string, 0), crates[instr.from-1][0:instr.nb]...)
		crateTo = append(crateTo, crates[instr.to-1]...)

		crates[instr.to-1] = crateTo
		crates[instr.from-1] = crateFrom
	}

	tops := make([]string, 0)
	for _, crate := range crates {
		tops = append(tops, crate[0])
	}

	return strings.Join(tops, "")
}

func LoadCrateData(fileName string) ([][]string, []instruction) {
	readFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var crates [][]string
	moves := make([]instruction, 0)

	for fileScanner.Scan() {
		currentLine := fileScanner.Text()

		st := strings.TrimSpace(currentLine)

		if len(st) == 0 {
			continue
		}

		switch st[0] {
		case craneStart[0]:
			nbStacks := (len(currentLine) + 1) / 4
			for i := 0; i < nbStacks; i++ {
				if len(crates) == 0 {
					crates = make([][]string, nbStacks)
				}
				cratePosition := (i * 4) + 1
				if crate := string(currentLine[cratePosition]); crate != " " {
					crates[i] = append(crates[i], crate)
				}
			}
		case instructionStart[0]:
			var (
				from int
				to   int
				nb   int
			)

			fmt.Sscanf(currentLine, "move %d from %d to %d", &nb, &from, &to)
			moves = append(moves, instruction{from: from, to: to, nb: nb})
		default:
		}
	}
	readFile.Close()

	return crates, moves
}

type instruction struct {
	from int
	to   int
	nb   int
}

const (
	craneStart       string = "["
	instructionStart string = "m"
)
