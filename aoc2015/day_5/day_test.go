package day_5

import (
	"advent-of-code/aoc/utils"
	"bufio"
	"os"
	"path"
	"testing"
)

func TestQuestion1(t *testing.T) {
	var (
		expected = "2"
		val      string
	)
	day := create("input_test.txt")
	val = day.Question1()
	t.Log("expected", expected, "got", val)
	if val != expected {
		t.Fail()
	}
}

func TestQuestion2(t *testing.T) {
	var (
		expected = "2"
		val      string
	)
	day := create("input_test2.txt")
	val = day.Question2()
	t.Log("expected", expected, "got", val)
	if val != expected {
		t.Fail()
	}
}

func createTestFiles(testDir string) {
	file := utils.LoadInput("input_test.txt")
	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		os.WriteFile(path.Join(testDir, line+".input"), []byte(line), 0644)
	}
}
