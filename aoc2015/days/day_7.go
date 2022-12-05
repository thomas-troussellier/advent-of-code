package days

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day7() {
	fmt.Println("Day 7")

	// load exercise data
	instructions := loadCircuitInstr("aoc2015/inputs/day_7.log")

	registry := make(map[string]uint16)

	fmt.Println("Q1. What signal is ultimately provided to wire a?")
	regInstr, err := regexp.Compile("[A-Z]+")
	if err != nil {
		fmt.Println(err)
	}

	doAssignments(registry, regInstr, instructions)

	fmt.Printf("A1. signal: a:%d\n", registry["a"])

	fmt.Println("Q2. take the signal you got on wire a, override wire b to that signal, and reset the other wires (including wire a). What new signal is ultimately provided to wire a?")
	valueOfa := registry["a"]
	for key := range registry {
		delete(registry, key)
	}
	for index, instruction := range instructions {
		if strings.HasSuffix(instruction, "-> b") {
			instructions[index] = fmt.Sprintf("%d -> b", valueOfa)
		}
	}

	doAssignments(registry, regInstr, instructions)
	fmt.Printf("A2. signal: a:%d\n", registry["a"])
}

func doAssignments(registry map[string]uint16, regInstr *regexp.Regexp, instructions []string) {
	followUp := make([]string, 0)
	for _, instruction := range instructions {
		if i := doAssignment(registry, regInstr, instruction); i != nil {
			followUp = append(followUp, instruction)
		}
	}

	if len(followUp) > 0 {
		doAssignments(registry, regInstr, followUp)
	}
}

func doAssignment(registry map[string]uint16, regInstr *regexp.Regexp, instruction string) error {

	rInst := regInstr.Find([]byte(instruction))
	var stringInst string
	if rInst == nil {
		stringInst = ASSIGN
	} else {
		stringInst = string(rInst)
	}

	switch stringInst {
	case NOT:
		var fromReg, toReg string
		_, err := fmt.Sscanf(instruction, "NOT %s -> %s", &fromReg, &toReg)
		if err != nil {
			fmt.Println(err)
		}

		uNot, err := convertToUint16(registry, fromReg)
		if err != nil {
			return ErrWait
		}
		registry[toReg] = ^uNot
	case ASSIGN:
		var (
			toReg  string
			signal string
		)
		_, err := fmt.Sscanf(instruction, "%s -> %s", &signal, &toReg)
		if err != nil {
			fmt.Println(err)
		}

		uAssign, err := convertToUint16(registry, signal)
		if err != nil {
			return ErrWait
		}
		registry[toReg] = uAssign
	case AND:
		var and1, and2, toReg string
		_, err := fmt.Sscanf(instruction, "%s AND %s -> %s", &and1, &and2, &toReg)
		if err != nil {
			fmt.Println(err)
		}
		var uAnd1, uAnd2 uint16
		uAnd1, err = convertToUint16(registry, and1)
		if err != nil {
			return ErrWait
		}
		uAnd2, err = convertToUint16(registry, and2)
		if err != nil {
			return ErrWait
		}
		registry[toReg] = uAnd1 & uAnd2
	case OR:
		var or1, or2, toReg string
		_, err := fmt.Sscanf(instruction, "%s OR %s -> %s", &or1, &or2, &toReg)
		if err != nil {
			fmt.Println(err)
		}
		var uor1, uor2 uint16
		uor1, err = convertToUint16(registry, or1)
		if err != nil {
			return ErrWait
		}
		uor2, err = convertToUint16(registry, or2)
		if err != nil {
			return ErrWait
		}

		registry[toReg] = uor1 | uor2
	case RSHIFT:
		var shift1, shift2, toReg string
		_, err := fmt.Sscanf(instruction, "%s RSHIFT %s -> %s", &shift1, &shift2, &toReg)
		if err != nil {
			fmt.Println(err)
		}
		var ushift1, ushift2 uint16
		ushift1, err = convertToUint16(registry, shift1)
		if err != nil {
			return ErrWait
		}
		ushift2, err = convertToUint16(registry, shift2)
		if err != nil {
			return ErrWait
		}

		registry[toReg] = ushift1 >> ushift2
	case LSHIFT:
		var shift1, shift2, toReg string
		_, err := fmt.Sscanf(instruction, "%s LSHIFT %s -> %s", &shift1, &shift2, &toReg)
		if err != nil {
			fmt.Println(err)
		}
		var ushift1, ushift2 uint16
		ushift1, err = convertToUint16(registry, shift1)
		if err != nil {
			return ErrWait
		}
		ushift2, err = convertToUint16(registry, shift2)
		if err != nil {
			return ErrWait
		}

		registry[toReg] = ushift1 << ushift2
	}
	return nil
}

func convertToUint16(registry map[string]uint16, input string) (uint16, error) {
	var res uint16
	if u, err := strconv.ParseUint(input, 10, 16); err == nil {
		res = uint16(u)
	} else {
		val, ok := registry[input]
		if !ok {
			return res, ErrWait
		}
		res = val
	}

	return res, nil
}

const (
	NOT    = "NOT"
	ASSIGN = "ASSIGN"
	AND    = "AND"
	OR     = "OR"
	RSHIFT = "RSHIFT"
	LSHIFT = "LSHIFT"
)

var ErrWait error = errors.New("waiting for value")

func loadCircuitInstr(fileName string) []string {
	readFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	list := make([]string, 0)

	for fileScanner.Scan() {
		instruction := fileScanner.Text()

		list = append(list, instruction)

	}

	readFile.Close()

	return list
}
