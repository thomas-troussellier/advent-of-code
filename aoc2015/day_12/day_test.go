package day_12

import (
	"os"
	"path"
	"strconv"
	"testing"
)

func TestQuestion1(t *testing.T) {
	testcases := []test{
		{
			value:    "[1,2,3]",
			expected: "6",
		},
		{
			value:    "{\"a\":2,\"b\":4}",
			expected: "6",
		},
		{
			value:    "[[[3]]]",
			expected: "3",
		},
		{
			value:    "{\"a\":{\"b\":4},\"c\":-1}",
			expected: "3",
		},
		{
			value:    "{\"a\":[-1,1]}",
			expected: "0",
		},
		{
			value:    "[-1,{\"a\":1}]",
			expected: "0",
		},
		{
			value:    "[]",
			expected: "0",
		},
		{
			value:    "{}",
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
			value:    "[1,2,3]",
			expected: "6",
		},
		{
			value:    "[1,{\"c\":\"red\",\"b\":2},3]",
			expected: "4",
		},
		{
			value:    "{\"d\":\"red\",\"e\":[1,2,3,4],\"f\":5}",
			expected: "0",
		},
		{
			value:    "[1,\"red\",5]",
			expected: "6",
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
