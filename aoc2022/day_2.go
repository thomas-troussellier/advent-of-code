package aoc2022

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func day2() {
	fmt.Println("Day 2")

	// load exercise data
	strategy := LoadRockPaperScissorsData("aoc2022/day_2.log")

	// based on prior assumption that X,Y,Z (rock, paper, scissors) correspond to your moves
	totalScore := 0
	for _, round := range strategy {
		round.Resolve()
		totalScore += round.score
	}
	fmt.Println("Q1. What would your total score be if everything goes exactly according to your strategy guide?")
	fmt.Printf("A1. Total score will be: %d\n", totalScore)

	// based on new assumption that X,Y,Z (loss, draw, win) correspond to the result of the round instead
	totalScore = 0
	for _, round := range strategy {
		round.Resolve2()
		totalScore += round.score
	}
	fmt.Println("Q2. What would your total score be if everything goes exactly according to your strategy guide?")
	fmt.Printf("A2. Total score will be: %d\n", totalScore)
}

const (
	elf_rock     = "A"
	elf_paper    = "B"
	elf_scissors = "C"

	cypher_X = "X"
	cypher_Y = "Y"
	cypher_Z = "Z"
)

var points_mapping = map[string]int{
	elf_rock:     1,
	elf_paper:    2,
	elf_scissors: 3,
}

var to_win_against_mapping = map[string]string{
	elf_rock:     elf_paper,
	elf_paper:    elf_scissors,
	elf_scissors: elf_rock,
}

var to_lose_against_mapping = map[string]string{
	elf_rock:     elf_scissors,
	elf_paper:    elf_rock,
	elf_scissors: elf_paper,
}

type round struct {
	enemyMove string
	cypher    string
	score     int
}

type strategy []*round

func (r *round) Resolve() {
	// your move for the round is cypher
	r.score = 0
	var move string
	switch r.cypher {
	case cypher_X: // rock
		r.score += points_mapping[elf_rock]
		move = elf_rock
	case cypher_Y: // paper
		r.score += points_mapping[elf_paper]
		move = elf_paper
	case cypher_Z: // scissors
		r.score += points_mapping[elf_scissors]
		move = elf_scissors
	}

	r.score += match(move, r.enemyMove)
}

func match(yourMove, enemyMove string) int {
	if move := to_win_against_mapping[enemyMove]; move == yourMove {
		return 6
	} else if move := to_lose_against_mapping[enemyMove]; move == yourMove {
		return 0
	}
	return 3
}

func (r *round) Resolve2() {
	// round outcome must correspond to cypher
	r.score = 0

	switch r.cypher {
	case cypher_X: // loss
		expectedMove := to_lose_against_mapping[r.enemyMove]
		r.score += points_mapping[expectedMove]
	case cypher_Y: // draw
		r.score += points_mapping[r.enemyMove]
		r.score += 3
	case cypher_Z: // win
		expectedMove := to_win_against_mapping[r.enemyMove]
		r.score += points_mapping[expectedMove]
		r.score += 6
	}
}

func LoadRockPaperScissorsData(fileName string) strategy {
	readFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	rounds := make([]*round, 0, 10)

	for fileScanner.Scan() {
		currentLine := fileScanner.Text()

		lineTokens := strings.Split(currentLine, " ")

		currentMove := &round{
			enemyMove: lineTokens[0],
			cypher:    lineTokens[1],
		}
		rounds = append(rounds, currentMove)
	}
	readFile.Close()

	return rounds
}
