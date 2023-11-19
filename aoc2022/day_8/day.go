package day_8

import (
	"advent-of-code/aoc"
	"bufio"
	"log"
	"os"
	"strconv"
)

type day struct {
	inputFile string
}

var _ aoc.EventRunner = (*day)(nil)

func Create() aoc.EventRunner {
	return create("aoc2022/day_8/input.txt")
}

func create(inputFile string) *day {
	return &day{
		inputFile: inputFile,
	}
}

func (d *day) Execute() {
	log.Println("Day 8")
	log.Println("Q1:", d.Question1())
	log.Println("Q2:", d.Question2())
}

func (d *day) Question1() string {
	trees := loadForest(d.inputFile, 4)
	return strconv.Itoa(computeVisibleTrees(trees))
}

func (d *day) Question2() string {
	trees := loadForest(d.inputFile, 4)
	return strconv.Itoa(computeScenicScore(trees))
}

func computeScenicScore(trees [][]int) int {
	highest := 0
	for i := 1; i < len(trees)-1; i++ {
		currentTreeLine := trees[i]
		for j := 1; j < len(currentTreeLine)-1; j++ {
			score := 0
			_, s := isVisible(currentTreeLine, j)
			score += s

			currentTreeColumn := treeColumn(trees, j)
			_, s = isVisible(currentTreeColumn, i)
			score *= s

			if score > highest {
				highest = score
			}
		}
	}
	return highest
}

func computeVisibleTrees(trees [][]int) int {
	exterior := ((len(trees[0]) - 2) * 2) + (len(trees) * 2)

	interior := 0
	for i := 1; i < len(trees)-1; i++ {
		currentTreeLine := trees[i]
		for j := 1; j < len(currentTreeLine)-1; j++ {
			if v, _ := isVisible(currentTreeLine, j); v {
				interior++
				continue
			}
			currentTreeColumn := treeColumn(trees, j)
			if v, _ := isVisible(currentTreeColumn, i); v {
				interior++
				continue
			}
		}
	}

	return exterior + interior
}

func treeColumn(trees [][]int, columnIndex int) (column []int) {
	column = make([]int, 0)
	for _, row := range trees {
		column = append(column, row[columnIndex])
	}
	return
}

func isVisible(treeLine []int, treeIndex int) (vis bool, scenicScore int) {
	if treeIndex == 0 || treeIndex == (len(treeLine)-1) {
		return true, 0
	}

	// need to reverse this
	left := reverse(treeLine[:treeIndex])
	vl, nbl := visible(left, treeLine[treeIndex])
	scenicScore += nbl

	right := treeLine[treeIndex+1:]
	vr, nbr := visible(right, treeLine[treeIndex])
	scenicScore *= nbr

	return (vl || vr), scenicScore
}

func reverse(numbers []int) []int {
	newNumbers := make([]int, len(numbers))
	for i, j := 0, len(numbers)-1; i <= j; i, j = i+1, j-1 {
		newNumbers[i], newNumbers[j] = numbers[j], numbers[i]
	}
	return newNumbers
}

func visible(treeLine []int, treesize int) (vis bool, nbTree int) {
	vis = true
	nbVisible := 0
	for i := 0; i < len(treeLine); i++ {
		nbVisible++
		if treeLine[i] >= treesize {
			vis = false
			break
		}
	}

	return vis, nbVisible
}

func loadForest(fileName string, diffChars int) [][]int {
	readFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	trees := make([][]int, 0)

	for fileScanner.Scan() {
		currentLine := fileScanner.Text()

		treeLine := make([]int, 0)
		for i := 0; i < len(currentLine); i++ {
			tree, _ := strconv.Atoi(string(currentLine[i]))
			treeLine = append(treeLine, tree)
		}
		trees = append(trees, treeLine)
	}

	readFile.Close()

	return trees
}
