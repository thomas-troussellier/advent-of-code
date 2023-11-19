package day_12

import (
	"advent-of-code/aoc"
	"advent-of-code/aoc/utils"
	"bufio"
	"encoding/json"
	"log"
	"strconv"
)

type day struct {
	inputFile string
}

var _ aoc.EventRunner = (*day)(nil)

func Create() aoc.EventRunner {
	return create("aoc2015/day_12/input.txt")
}

func create(inputFile string) *day {
	return &day{
		inputFile: inputFile,
	}
}

func (d *day) Execute() {
	log.Println("Day 12")
	log.Println("Q1:", d.Question1())
	log.Println("Q2:", d.Question2())
}

func (d *day) Question1() string {
	jsonString := loadJson(d.inputFile)
	json := parseJson(jsonString)
	return strconv.FormatFloat(computeSum(json, false), 'f', -1, 64)
}

func (d *day) Question2() string {
	jsonString := loadJson(d.inputFile)
	json := parseJson(jsonString)
	return strconv.FormatFloat(computeSum(json, true), 'f', -1, 64)
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
	readFile := utils.LoadInput(fileName)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var jsonString string
	for fileScanner.Scan() {
		jsonString = fileScanner.Text()
	}

	readFile.Close()

	return jsonString
}
