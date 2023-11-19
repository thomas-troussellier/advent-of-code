package utils

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func GetInputFromSite(year int, day int) {
	inputFileName := fmt.Sprintf("aoc%d/day_%d/input.txt", year, day)
	fileInfo, err := os.Stat(inputFileName)
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		log.Fatal("an error occured", err)
	}
	if fileInfo != nil {
		log.Println("input file already exists for", year, day)
		return
	}

	cookie, ok := os.LookupEnv("AOC_SESSION_COOKIE")
	if !ok {
		log.Fatal("missing env var AOC_SESSION_COOKIE")
	}

	if cookie == "" {
		log.Fatal("your AOC_SESSION_COOKIE is empty")
	}

	aocCookie := &http.Cookie{Value: cookie, Name: "session"}

	aocInputUrl := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	getRequest, err := http.NewRequest(http.MethodGet, aocInputUrl, nil)
	if err != nil {
		log.Fatal("failed to build getInput request", err)
	}
	getRequest.AddCookie(aocCookie)

	resp, err := http.DefaultClient.Do(getRequest)
	if err != nil || resp.StatusCode != 200 {
		log.Fatal("error while getting input", err)
	}
	defer resp.Body.Close()

	os.MkdirAll(filepath.Dir(inputFileName), os.ModePerm)
	inputFile, err := os.Create(inputFileName)
	if err != nil {
		log.Fatal("error creating input file", err)
	}
	defer inputFile.Close()
	_, err = inputFile.ReadFrom(resp.Body)
	if err != nil {
		log.Fatal("failed to write request body to input file", err)
	}
	log.Println("got input file for", year, day)
}
