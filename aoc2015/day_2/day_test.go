package day_2

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
			value:    "2x3x4",
			expected: "58",
		},
		{
			name:     "b",
			value:    "1x1x10",
			expected: "43",
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
			value:    "2x3x4",
			expected: "34",
		},
		{
			name:     "b",
			value:    "1x1x10",
			expected: "14",
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
