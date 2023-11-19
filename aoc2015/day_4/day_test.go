package day_4

import (
	"os"
	"path"
	"strconv"
	"testing"
)

func TestQuestion1(t *testing.T) {
	testcases := []test{
		{
			value:    "abcdef",
			expected: "609043",
		},
		{
			value:    "pqrstuv",
			expected: "1048970",
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

type test struct {
	value    string
	expected string
}

func createTestFile(testDir string, value string, name int) string {
	fileName := path.Join(testDir, strconv.Itoa(name)+".input")
	os.WriteFile(fileName, []byte(value), 0644)
	return fileName
}
