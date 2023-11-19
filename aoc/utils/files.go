package utils

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strconv"
	"strings"
)

func LoadInput(inputFileName string) *os.File {
	readFile, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal("failed to open input file", err)
	}

	return readFile
}

func CopyTemplates(year, day int) {
	inputBase := "aoc/templates/"
	files, err := os.ReadDir(inputBase)
	if err != nil {
		log.Fatal("failed to copy template files. ", err)
	}

	destinationFileBase := fmt.Sprintf("aoc%d/day_%d/", year, day)

	for _, file := range files {
		extName, _ := strings.CutSuffix(file.Name(), ".template")
		destinationName := destinationFileBase + extName

		if info, err := os.Stat(destinationName); err == nil {
			log.Println("file", info.Name(), "already exists")
			continue
		} else if !errors.Is(err, fs.ErrNotExist) {
			log.Fatal("error while checking file system. ", err)
		}

		input, err := os.ReadFile(inputBase + file.Name())
		if err != nil {
			log.Fatal("failed to read template file. ", err)
		}
		if strings.HasPrefix(extName, "day") {
			input = bytes.ReplaceAll(input, []byte("DAY_PLACEHOLDER"), []byte(strconv.Itoa(day)))
			input = bytes.ReplaceAll(input, []byte("YEAR_PLACEHOLDER"), []byte(strconv.Itoa(year)))
		}
		err = os.WriteFile(destinationName, input, 0644)
		if err != nil {
			log.Fatal("failed to copy template. ", err)
		}
	}

	log.Printf("templates files for year %d day %d created", year, day)
}
