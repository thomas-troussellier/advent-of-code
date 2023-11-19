package day_1

import (
	"os"
	"path"
	"testing"
)

type test struct {
	name     string
	value    string
	expected string
}

func TestQuestion1(t *testing.T) {
	testcases := []test{
		{
			name:     "a",
			value:    "(())",
			expected: "0",
		},
		{
			name:     "b",
			value:    "()()",
			expected: "0",
		},
		{
			name:     "c",
			value:    "(((",
			expected: "3",
		},
		{
			name:     "d",
			value:    "(()(()(",
			expected: "3",
		},
		{
			name:     "e",
			value:    "))(((((",
			expected: "3",
		},
		{
			name:     "f",
			value:    "())",
			expected: "-1",
		},
		{
			name:     "g",
			value:    "))(",
			expected: "-1",
		},
		{
			name:     "h",
			value:    ")))",
			expected: "-3",
		},
		{
			name:     "i",
			value:    ")())())",
			expected: "-3",
		},
	}

	testDir := t.TempDir()
	for _, test := range testcases {
		day := create(createTestFile(testDir, test))
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
			name:     "a",
			value:    ")",
			expected: "1",
		},
		{
			name:     "b",
			value:    "()())",
			expected: "5",
		},
	}

	testDir := t.TempDir()
	for _, test := range testcases {
		day := create(createTestFile(testDir, test))
		val := day.Question2()
		t.Log("expected", test.expected, "got", val)
		if val != test.expected {
			t.Fail()
		}
	}
}

func createTestFile(testDir string, t test) string {
	fileName := path.Join(testDir, t.name+".input")
	os.WriteFile(fileName, []byte(t.value), 0644)
	return fileName
}
