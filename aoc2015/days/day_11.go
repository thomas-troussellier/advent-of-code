package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Day11() {
	fmt.Println("Day 11")

	// load exercise data
	password := loadPassword("aoc2015/inputs/day_11.log")

	fmt.Println("Q1. What should his next password be?")
	newPass := computeNextPassword(password, false)
	fmt.Printf("A1. Next password: %s\n", newPass)

	fmt.Println("Q2. Santa's password expired again. What's the next one?")
	newPass = computeNextPassword(newPass, true)
	fmt.Printf("A2. New password: %s\n", newPass)
}

func computeNextPassword(password string, forceRegen bool) string {
	valid, allowed, fIndex := checkPassRules(password)

	if valid && !forceRegen {
		return password
	}

	nextPass := password

	// make replace first forbidden char from pass
	if !allowed {
		nextPass = strings.Replace(nextPass, string(nextPass[fIndex]), string(nextPass[fIndex]+1), 1)
		nextPass = nextPass[:fIndex+1] + strings.Repeat("a", len(nextPass[fIndex+1:]))
		return computeNextPassword(nextPass, false)
	}

	nextPass = getNextIncPass(nextPass)
	return computeNextPassword(nextPass, false)
}

func getNextIncPass(input string) string {
	endChar := input[len(input)-1]

	if endChar == 'z' {
		return getNextIncPass(input[:len(input)-1]) + "a"
	}

	return input[:len(input)-1] + string(endChar+1)
}

func checkPassRules(pass string) (isValid, allowed bool, fIndex int) {
	var length, seq bool
	var dupCount int

	if len(pass) == 8 {
		length = true
	}

	if strings.ContainsAny(pass, "iol") {
		fIndex = strings.IndexAny(pass, "iol")
	} else {
		allowed = true
	}

	for i := 0; i < len(pass)-2; i++ {
		if pass[i] == pass[i+1]-1 && pass[i+1] == pass[i+2]-1 {
			seq = true
			break
		}
	}

	for i := 0; i < len(pass)-1; i++ {
		if pass[i] == pass[i+1] {
			dupCount++
			i += 1
		}
		if dupCount == 2 {
			break
		}
	}

	isValid = length && allowed && seq && (dupCount == 2)

	return
}

func loadPassword(fileName string) string {
	readFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var password string
	for fileScanner.Scan() {
		password = fileScanner.Text()
	}

	readFile.Close()

	return password
}
