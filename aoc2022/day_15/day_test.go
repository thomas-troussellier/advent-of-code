package day_15

import (
	"strconv"
	"testing"
)

func TestQuestion1(t *testing.T) {
	var (
		expected = "26"
		val      string
	)
	day := create("input_test.txt")
	s, b := day.loadData()
	val = strconv.Itoa(nbUnavailableSpacesOnLine(s, b, 10))
	t.Log("expected", expected, "got", val)
	if val != expected {
		t.Fail()
	}
}

func TestQuestion2(t *testing.T) {
	var (
		expected = "56000011"
		val      string
	)
	day := create("input_test.txt")
	s, b := day.loadData()
	val = strconv.Itoa(findFrequency(s, b, 20))
	t.Log("expected", expected, "got", val)
	if val != expected {
		t.Fail()
	}
}
