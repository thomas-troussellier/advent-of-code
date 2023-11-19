package day_4

import (
	"advent-of-code/aoc"
	"advent-of-code/aoc/utils"
	"bufio"
	"bytes"
	"crypto/md5"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type day struct {
	inputFile string
}

var _ aoc.EventRunner = (*day)(nil)

func Create() aoc.EventRunner {
	return create("aoc2015/day_4/input.txt")
}

func create(inputFile string) *day {
	return &day{
		inputFile: inputFile,
	}
}

func (d *day) Execute() {
	log.Println("Day 4")
	log.Println("Q1:", d.Question1())
	log.Println("Q2:", d.Question2())
}

func (d *day) Question1() string {
	secret := loadSecret(d.inputFile)
	cmpt := 1
	for {
		temp := secret + strconv.Itoa(cmpt)
		solution := md5.Sum([]byte(temp))
		if bytes.HasPrefix(solution[:], nilBits) && strings.HasPrefix(fmt.Sprintf("%x", solution), "00000") {
			break
		}
		cmpt++
	}

	return strconv.Itoa(cmpt)
}

func (d *day) Question2() string {
	secret := loadSecret(d.inputFile)
	cmpt := 1
	for {
		temp := secret + strconv.Itoa(cmpt)
		solution := md5.Sum([]byte(temp))
		if bytes.HasPrefix(solution[:], moarNilBits) {
			break
		}
		cmpt++
	}

	return strconv.Itoa(cmpt)
}

var nilBits = []byte{0, 0}

var moarNilBits = []byte{0, 0, 0}

func loadSecret(fileName string) string {
	readFile := utils.LoadInput(fileName)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var secret string
	for fileScanner.Scan() {
		secret = fileScanner.Text()
	}

	readFile.Close()

	return secret
}
