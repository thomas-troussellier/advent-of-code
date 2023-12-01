package day_1

import "testing"

func TestQuestion1(t *testing.T) {
	var (
		expected = "142"
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
		expected = "281"
		val      string
	)
	day := create("input_test2.txt")
	val = day.Question2()
	t.Log("expected", expected, "got", val)
	if val != expected {
		t.Fail()
	}
}
