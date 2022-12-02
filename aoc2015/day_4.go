package aoc2015

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func day4() {
	fmt.Println("Day 4")

	// load exercise data
	secret := loadSecret("aoc2015/day_4.log")

	fmt.Println("Q1. need to find MD5 hashes which, in hexadecimal, start with at least five zeroes")
	cmpt := 1
	for {
		temp := secret + strconv.Itoa(cmpt)
		solution := md5.Sum([]byte(temp))
		if bytes.HasPrefix(solution[:], nilBits) && strings.HasPrefix(fmt.Sprintf("%x", solution), "00000") {
			break
		}
		cmpt++
	}
	fmt.Printf("A1. solution : %s, md5: %x\n", secret+strconv.Itoa(cmpt), md5.Sum([]byte(secret+strconv.Itoa(cmpt))))

	fmt.Println("Q2. Now find one that starts with six zeroes")
	cmpt = 1
	for {
		temp := secret + strconv.Itoa(cmpt)
		solution := md5.Sum([]byte(temp))
		if bytes.HasPrefix(solution[:], moarNilBits) {
			break
		}
		cmpt++
	}
	fmt.Printf("A1. solution : %s, md5: %x\n", secret+strconv.Itoa(cmpt), md5.Sum([]byte(secret+strconv.Itoa(cmpt))))

}

var nilBits = []byte{0, 0}

var moarNilBits = []byte{0, 0, 0}

func loadSecret(fileName string) string {
	readFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var secret string
	for fileScanner.Scan() {
		secret = fileScanner.Text()
	}

	readFile.Close()

	return secret
}
