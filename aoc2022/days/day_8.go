package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Day8() {
	fmt.Println("Day 8")
	trees := loadForest("aoc2022/inputs/day_8.log", 4)

	fmt.Println("Q1. How many trees are visible from outside the grid?")
	visibleTrees := computeVisibleTrees(trees)
	fmt.Printf("A1. visible trees: %d\n", visibleTrees)

	fmt.Println("Q2. What is the highest scenic score possible for any tree?")
	scenicScore := computeScenicScore(trees)
	fmt.Printf("A2. Highest scenic score: %d\n", scenicScore)
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
