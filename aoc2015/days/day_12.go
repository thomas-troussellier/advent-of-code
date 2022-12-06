package days

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func Day12() {
	fmt.Println("Day 12")

	// load exercise data
	jsonString := loadJson("aoc2015/inputs/day_12.log")
	json := parseJson(jsonString)
	fmt.Println("Q1. What is the sum of all numbers in the document?")
	sum := computeSum(json, false)
	fmt.Printf("A1. Sum of all numbers: %d\n", int(sum))

	fmt.Println("Q2. What is the sum of all numbers in the document?")
	sum = computeSum(json, true)
	fmt.Printf("A2. Sum of all numbers: %d\n", int(sum))
}

func parseJson(input string) interface{} {
	data := []byte(input)
	var obj interface{}
	if err := json.Unmarshal(data, &obj); err != nil {
		log.Fatal(err)
	}

	return obj
}

func computeSum(json interface{}, ignoreRed bool) float64 {
	tot, _ := computeSumofAllNumbers(json, ignoreRed)
	return tot
}

func computeSumofAllNumbers(json interface{}, ignoreRed bool) (sum float64, red bool) {

switchLabel:
	switch v := json.(type) {
	case []interface{}:
		for _, elmt := range v {
			tot, _ := computeSumofAllNumbers(elmt, ignoreRed)
			sum += tot
		}
	case map[string]interface{}:
		var tmp float64 = 0
		for _, elmt := range v {
			total, isRed := computeSumofAllNumbers(elmt, ignoreRed)
			tmp += total
			if ignoreRed && isRed {
				break switchLabel
			}
		}
		sum += tmp
	case float64:
		sum += v
	case string:
		if v == "red" {
			return sum, true
		}
	default:
	}

	return
}

func loadJson(fileName string) string {
	readFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var jsonString string
	for fileScanner.Scan() {
		jsonString = fileScanner.Text()
	}

	readFile.Close()

	return jsonString
}
