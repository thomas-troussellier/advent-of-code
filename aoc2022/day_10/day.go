package day_10

import (
	"advent-of-code/aoc"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type day struct {
	inputFile string
}

var _ aoc.EventRunner = (*day)(nil)

func Create() aoc.EventRunner {
	return create("aoc2022/day_10/input.txt")
}

func create(inputFile string) *day {
	return &day{
		inputFile: inputFile,
	}
}

func (d *day) Execute() {
	log.Println("Day 10")
	log.Println("Q1:", d.Question1())
	log.Println("Q2:", d.Question2())
}

func (d *day) Question1() string {
	program := loadProgram(d.inputFile)

	signalPerCycles := executeProgram(program)

	return strconv.Itoa(computeStrengthForCycles(signalPerCycles, 20, 60, 100, 140, 180, 220))
}

func (d *day) Question2() string {
	program := loadProgram(d.inputFile)

	signalPerCycles := executeProgram(program)
	printCRTImage(signalPerCycles)

	return "Q2"
}

func printCRTImage(signalPerCycles map[int]int) {
	for i := 1; i <= 6; i++ {
		printCRTLine(signalPerCycles, i)
	}
}

func printCRTLine(signalPerCycles map[int]int, line int) {
	var builder strings.Builder

	for column := 1; column <= 40; column++ {
		var pixel rune
		currentlyDrawingPixel := column
		if pixelSeesSprite(currentlyDrawingPixel, signalPerCycles[((line-1)*40)+column]+1) {
			pixel = '#'
		} else {
			pixel = '.'
		}
		builder.WriteRune(pixel)
	}
	fmt.Println(builder.String())
}

func pixelSeesSprite(pixel, spriteMiddle int) bool {
	if pixel == spriteMiddle {
		return true
	}
	if pixel == spriteMiddle+1 || pixel == spriteMiddle-1 {
		return true
	}
	return false
}

func computeStrengthForCycles(signalPerCycles map[int]int, cycles ...int) int {
	total := 0
	for _, cycle := range cycles {
		fmt.Println(cycle, signalPerCycles[cycle])
		total += signalPerCycles[cycle] * cycle
	}
	return total
}

func executeProgram(program []progInstr) map[int]int {
	cycles := make(map[int]int)
	cycles[0] = 1
	cycle := 0
	strength := 1
	for _, instr := range program {
		switch instr.name {
		case "noop":
			cycle++
			cycles[cycle] = strength
		default:
			cycle++
			cycles[cycle] = strength
			cycle++
			cycles[cycle] = strength
			strength += instr.value
		}
	}
	cycle++
	cycles[cycle] = strength

	return cycles
}

type cycleState struct {
	sprite int
	signal int
}

type progInstr struct {
	name  string
	value int
}

func loadProgram(fileName string) []progInstr {
	readFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	prog := make([]progInstr, 0)

	for fileScanner.Scan() {
		currentLine := fileScanner.Text()

		instrLine := strings.Split(currentLine, " ")

		i := progInstr{
			name: instrLine[0],
		}
		if len(instrLine) == 2 {
			i.value, _ = strconv.Atoi(instrLine[1])
		}

		prog = append(prog, i)
	}

	readFile.Close()

	return prog
}
