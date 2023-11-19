package day_13

import "testing"

func TestQuestion1(t *testing.T) {
	var (
		expected = "330"
		val      string
	)
	day := create("input_test.txt")
	val = day.Question1()
	t.Log("expected", expected, "got", val)
	if val != expected {
		t.Fail()
	}
}
