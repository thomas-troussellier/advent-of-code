package day_11

import "testing"

func TestQuestion1(t *testing.T) {
	var (
		expected = "10605"
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
		expected = "2713310158"
		val      string
	)
	day := create("input_test.txt")
	val = day.Question2()
	t.Log("expected", expected, "got", val)
	if val != expected {
		t.Fail()
	}
}
