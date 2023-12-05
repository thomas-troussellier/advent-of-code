package day_5

import (
	"advent-of-code/aoc"
	"advent-of-code/aoc/utils"
	"bufio"
	"log"
	"strconv"
	"strings"
)

type day struct {
	inputFile string
}

var _ aoc.EventRunner = (*day)(nil)

func Create() aoc.EventRunner {
	return create("aoc2023/day_5/input.txt")
}

func create(inputFile string) *day {
	return &day{
		inputFile: inputFile,
	}
}

func (d *day) Execute() {
	log.Println("Day 5")
	log.Println("Q1:", d.Question1())
	log.Println("Q2:", d.Question2())
}

func (d *day) Question1() string {
	al := d.loadData()

	min := 0
	for _, seed := range al.seeds {
		i := computeLocForSeed(al, seed)
		if (min == 0) || (i < min) {
			min = i
		}
	}

	return strconv.Itoa(min)
}

func (d *day) Question2() string {
	al := d.loadData()

	min := 0
	for i := 0; i < len(al.seeds)-1; i += 2 {
		start, _ := strconv.Atoi(al.seeds[i])
		size, _ := strconv.Atoi(al.seeds[i+1])
		end := start + size - 1
		loc := computeLocForSeedRange(al, start, end)
		if (min == 0) || (loc < min) {
			min = loc
		}
		log.Println(i, loc, min)
	}

	return strconv.Itoa(min)
}

func computeLocForSeed(al almanach, seed string) int {
	dist := seed
	dist = computeDest(al.seedsToSoil, dist)
	dist = computeDest(al.soilToFertilizer, dist)
	dist = computeDest(al.fertilizerToWater, dist)
	dist = computeDest(al.waterToLight, dist)
	dist = computeDest(al.lightToTemperature, dist)
	dist = computeDest(al.temperatureToHumidity, dist)
	dist = computeDest(al.humidityToLocation, dist)

	i, _ := strconv.Atoi(dist)
	return i
}

func computeLocForSeedRange(al almanach, seedStart, seedEnd int) int {
	min := 0

	for s := seedStart; s <= seedEnd; s++ {
		i := computeLocForSeed(al, strconv.Itoa(s))
		if (min == 0) || (i < min) {
			min = i
		}
	}

	return min
}

func splitWork(al almanach, seedStart, seedEnd int) int {

	return 0
}

type almanach struct {
	seedsToSoil           []string
	soilToFertilizer      []string
	fertilizerToWater     []string
	waterToLight          []string
	lightToTemperature    []string
	temperatureToHumidity []string
	humidityToLocation    []string
	seeds                 []string
}

func computeDest(corrMap []string, source string) string {
	s, _ := strconv.Atoi(source)
	loc := s
	for _, r := range corrMap {
		ranges := strings.Split(r, " ")
		sourceStart, _ := strconv.Atoi(ranges[1])
		rangeSize, _ := strconv.Atoi(ranges[2])
		sourceEnd := sourceStart + rangeSize - 1

		if (s < sourceStart) || (s > sourceEnd) {
			continue
		}

		destStart, _ := strconv.Atoi(ranges[0])
		loc = destStart - sourceStart + s
	}

	return strconv.Itoa(loc)
}

func (d *day) loadData() almanach {
	input := utils.LoadInput(d.inputFile)
	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)

	almanach := almanach{
		seeds:                 make([]string, 0),
		seedsToSoil:           make([]string, 0),
		soilToFertilizer:      make([]string, 0),
		fertilizerToWater:     make([]string, 0),
		waterToLight:          make([]string, 0),
		lightToTemperature:    make([]string, 0),
		temperatureToHumidity: make([]string, 0),
		humidityToLocation:    make([]string, 0),
	}
	mapsToUse := make([]*[]string, 0)
	mapsToUse = append(mapsToUse, &(almanach.seedsToSoil), &(almanach.soilToFertilizer), &(almanach.fertilizerToWater), &(almanach.waterToLight), &(almanach.lightToTemperature), &(almanach.temperatureToHumidity), &(almanach.humidityToLocation))
	ind := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()

		if strings.HasPrefix(line, "seeds:") {
			seeds, _ := strings.CutPrefix(line, "seeds: ")
			almanach.seeds = append(almanach.seeds, strings.Split(seeds, " ")...)
			continue
		}
		if line == "" {
			ind++
			continue
		}
		if strings.Contains(line, "map") {
			continue
		}
		*mapsToUse[ind-1] = append(*mapsToUse[ind-1], line)
	}
	return almanach
}
