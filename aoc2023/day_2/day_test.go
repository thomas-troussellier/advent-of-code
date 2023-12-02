package day_2

import (
	"strconv"
	"testing"
)

func TestQuestion1(t *testing.T) {
	var (
		expected = "8"
		val      string
	)
	day := create("input_test.txt")
	games := day.loadData()
	val = strconv.Itoa(day.sumIds(games, 12, 13, 14))
	t.Log("expected", expected, "got", val)
	if val != expected {
		t.Fail()
	}
}

func TestQuestion2(t *testing.T) {
	var (
		expected = "2286"
		val      string
	)
	day := create("input_test.txt")
	games := day.loadData()

	val = strconv.Itoa(day.computeTotalPower(games))
	t.Log("expected", expected, "got", val)
	if val != expected {
		t.Fail()
	}
}
