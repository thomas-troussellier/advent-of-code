package day_5

import (
	"advent-of-code/aoc"
	"advent-of-code/aoc/utils"
	"bufio"
	"fmt"
	"log"
	"strings"
)

type day struct {
	inputFile string
}

var _ aoc.EventRunner = (*day)(nil)

func Create() aoc.EventRunner {
	return create("aoc2022/day_5/input.txt")
}

func create(inputFile string) *day {
	return &day{
		inputFile: inputFile,
	}
}

func (d *day) Execute() {
	log.Println("Day 5")
	log.Println("Q1:", d.Question1())
	log.Println("Q2:", d.Question2())
}

func (d *day) Question1() string {
	crates, moves := loadCrateData(d.inputFile)

	return computeMoves9000(crates, moves)
}

func (d *day) Question2() string {
	crates, moves := loadCrateData(d.inputFile)
	return computeMoves9001(crates, moves)
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

func loadCrateData(fileName string) ([][]string, []instruction) {
	readFile := utils.LoadInput(fileName)

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
		case craneStart:
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
		case instructionStart:
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
	craneStart       byte = '['
	instructionStart byte = 'm'
)
