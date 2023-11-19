package day_6

import (
	"os"
	"path"
	"strconv"
	"testing"
)

func TestQuestion1(t *testing.T) {
	testcases := []test{
		{
			value:    "turn on 0,0 through 999,999",
			expected: "1000000",
		},
		{
			value:    "toggle 0,0 through 999,0",
			expected: "1000",
		},
		{
			value:    "turn off 499,499 through 500,500",
			expected: "0",
		},
	}

	testDir := t.TempDir()
	for i, test := range testcases {
		day := create(createTestFile(testDir, test.value, i))
		val := day.Question1()
		t.Log("expected", test.expected, "got", val)
		if val != test.expected {
			t.Fail()
		}
	}
}

func TestQuestion2(t *testing.T) {
	testcases := []test{
		{
			value:    "turn on 0,0 through 0,0",
			expected: "1",
		},
		{
			value:    "toggle 0,0 through 999,999",
			expected: "2000000",
		},
	}

	testDir := t.TempDir()
	for i, test := range testcases {
		day := create(createTestFile(testDir, test.value, i))
		val := day.Question2()
		t.Log("expected", test.expected, "got", val)
		if val != test.expected {
			t.Fail()
		}
	}
}

type test struct {
	value    string
	expected string
}

func createTestFile(testDir string, value string, name int) string {
	fileName := path.Join(testDir, strconv.Itoa(name)+".input")
	os.WriteFile(fileName, []byte(value), 0644)
	return fileName
}
