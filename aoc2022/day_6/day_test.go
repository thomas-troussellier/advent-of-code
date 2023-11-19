package day_6

import (
	"advent-of-code/aoc/utils"
	"bufio"
	"os"
	"path"
	"testing"
)

type test struct {
	expected string
	file     string
}

func TestQuestion1(t *testing.T) {
	testCases := []test{
		{
			expected: "7",
			file:     "mjqjpqmgbljsphdztnvjfqwrcgsmlb.input",
		},
		{
			expected: "5",
			file:     "bvwbjplbgvbhsrlpgdmjqwftvncz.input",
		},
		{
			expected: "6",
			file:     "nppdvjthqldpwncqszvftbrmjlhg.input",
		},
		{
			expected: "10",
			file:     "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg.input",
		},
		{
			expected: "11",
			file:     "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw.input",
		},
	}

	testDir := t.TempDir()
	createTestFiles(testDir)

	for _, test := range testCases {
		t.Run(test.file, func(t *testing.T) {
			day := create(path.Join(testDir, test.file))
			val := day.Question1()
			t.Log("expected", test.expected, "got", val)
			if val != test.expected {
				t.Fail()
			}
		})
	}
}

func TestQuestion2(t *testing.T) {
	testCases := []test{
		{
			expected: "19",
			file:     "mjqjpqmgbljsphdztnvjfqwrcgsmlb.input",
		},
		{
			expected: "23",
			file:     "bvwbjplbgvbhsrlpgdmjqwftvncz.input",
		},
		{
			expected: "23",
			file:     "nppdvjthqldpwncqszvftbrmjlhg.input",
		},
		{
			expected: "29",
			file:     "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg.input",
		},
		{
			expected: "26",
			file:     "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw.input",
		},
	}

	testDir := t.TempDir()
	createTestFiles(testDir)

	for _, test := range testCases {
		t.Run(test.file, func(t *testing.T) {
			day := create(path.Join(testDir, test.file))
			val := day.Question2()
			t.Log("expected", test.expected, "got", val)
			if val != test.expected {
				t.Fail()
			}
		})
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
