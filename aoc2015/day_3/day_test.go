package day_3

import (
	"os"
	"path"
	"strconv"
	"testing"
)

type test struct {
	value    string
	expected string
}

func TestQuestion1(t *testing.T) {
	testcases := []test{
		{
			value:    ">",
			expected: "2",
		},
		{
			value:    "^>v<",
			expected: "4",
		},
		{
			value:    "^v^v^v^v^v",
			expected: "2",
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
			value:    "^v",
			expected: "3",
		},
		{
			value:    "^>v<",
			expected: "3",
		},
		{
			value:    "^v^v^v^v^v",
			expected: "11",
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

func createTestFile(testDir string, value string, name int) string {
	fileName := path.Join(testDir, strconv.Itoa(name)+".input")
	os.WriteFile(fileName, []byte(value), 0644)
	return fileName
}
