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

	// load exercise data
	lookSay := loadLookSay("aoc2015/inputs/day_10.log")

	fmt.Println("Q1. What is the length of the result? (Apply 40 times)")
	res := computeLookSay(lookSay, 40)
	fmt.Printf("A1.Length of the result: %d\n", len(res))

	fmt.Println("Q1. What is the length of the result? (Apply 50 times)")
	res = computeLookSay(lookSay, 50)
	fmt.Printf("A1.Length of the result: %d\n", len(res))
}

func computeLookSay(input string, repeat int) string {

	tempString := input
	res := make([]string, 0)

	for i := 0; i < repeat; i++ {
		for len(tempString) > 0 {
			newInput := strings.TrimLeft(tempString, string(tempString[0]))
			occ := len(tempString) - len(newInput)
			res = append(res, strings.Join([]string{strconv.Itoa(occ), string(tempString[0])}, ""))
			tempString = newInput
		}
		tempString = strings.Join(res, "")
		res = make([]string, 0)
	}

	return tempString
}

func loadLookSay(fileName string) string {
	readFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var lookSayString string
	for fileScanner.Scan() {
		lookSayString = fileScanner.Text()
	}

	readFile.Close()

	return lookSayString
}
