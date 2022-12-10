package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day10() {
	fmt.Println("Day 10")

	program := loadProgram("aoc2022/inputs/day_10.log")

	signalPerCycles := executeProgram(program)

	fmt.Println("Q1. Find the signal strength during the 20th, 60th, 100th, 140th, 180th, and 220th cycles. What is the sum of these six signal strengths?")
	signalStrength := computeStrengthForCycles(signalPerCycles, 20, 60, 100, 140, 180, 220)
	fmt.Printf("A1. sum of these six signal strengths: %d\n", signalStrength)

	//slice := make(sort.IntSlice, 0)
	//for key := range signalPerCycles {
	//	slice = append(slice, key)
	//}
	//sort.Sort(slice)
	//for key := range slice {
	//	fmt.Println(key, signalPerCycles[key])
	//}

	fmt.Println("Q2. Render the image given by your program. What eight capital letters appear on your CRT?")
	printCRTImage(signalPerCycles)
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
