package day_5

import (
	"log"
	"testing"
)

func TestQuestion1(t *testing.T) {
	var (
		expected = "35"
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
		expected = "46"
		val      string
	)
	day := create("input_test.txt")
	val = day.Question2()
	t.Log("expected", expected, "got", val)
	if val != expected {
		t.Fail()
	}
}
func TestXxx(t *testing.T) {

	day := create("input_test.txt")
	al := day.loadData()

	m := make(map[int]int)
	for i := 0; i <= 100; i++ {
		m[i] = computeLocForSeed(al, i)
	}
	for i := 0; i <= 100; i++ {
		log.Println(i, m[i])
	}

	test := consolidate(al.seedsToSoil, al.soilToFertilizer)
	log.Println("al.seedsToSoil, al.soilToFertilizer", test)
	//test = consolidate(test, al.fertilizerToWater)
	//log.Println("fertilizerToWater", test)
	//test = consolidate(test, al.waterToLight)
	//log.Println("waterToLight", test)
	//test = consolidate(test, al.lightToTemperature)
	//log.Println("lightToTemperature", test)
	//test = consolidate(test, al.temperatureToHumidity)
	//log.Println("temperatureToHumidity", test)
	//test = consolidate(test, al.humidityToLocation)
	//log.Println("humidityToLocation", test)
}
