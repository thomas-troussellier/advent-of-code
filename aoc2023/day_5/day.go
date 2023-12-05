package day_5

import (
	"advent-of-code/aoc"
	"advent-of-code/aoc/utils"
	"bufio"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
	"sync"
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
		s, _ := strconv.Atoi(seed)
		i := computeLocForSeed(al, s)
		if (min == 0) || (i < min) {
			min = i
		}
	}

	return strconv.Itoa(min)
}

func (d *day) Question2() string {
	al := d.loadData()

	min := 0
	res := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < len(al.seeds)-1; i += 2 {
		wg.Add(1)
		go func(i int) {
			start, _ := strconv.Atoi(al.seeds[i])
			size, _ := strconv.Atoi(al.seeds[i+1])
			end := start + size - 1
			res <- computeLocForSeedRange(al, start, end)
			wg.Done()
		}(i)
	}
	go func() {
		wg.Wait()
		close(res)
	}()

	for r := range res {
		if (min == 0) || (r < min) {
			min = r
		}
	}

	return strconv.Itoa(min)
}

func computeLocForSeed(al almanach, seed int) int {
	dist := seed
	dist = computeDest(al.seedsToSoil, dist)
	dist = computeDest(al.soilToFertilizer, dist)
	dist = computeDest(al.fertilizerToWater, dist)
	dist = computeDest(al.waterToLight, dist)
	dist = computeDest(al.lightToTemperature, dist)
	dist = computeDest(al.temperatureToHumidity, dist)
	dist = computeDest(al.humidityToLocation, dist)

	return dist
}

func computeLocForSeedRange(al almanach, seedStart, seedEnd int) int {
	min := 0

	for s := seedStart; s <= seedEnd; s++ {
		i := computeLocForSeed(al, s)
		if (min == 0) || (i < min) {
			min = i
		}
	}

	return min
}

type almanach struct {
	seedsToSoil           []rang3
	soilToFertilizer      []rang3
	fertilizerToWater     []rang3
	waterToLight          []rang3
	lightToTemperature    []rang3
	temperatureToHumidity []rang3
	humidityToLocation    []rang3
	seeds                 []string
}

type rang3 struct {
	sourceStart, destStart, size int
}

var sortRange = func(a rang3, b rang3) int {
	return a.sourceStart - b.sourceStart
}

func rangeFromString(line string) rang3 {
	r := rang3{}
	fmt.Sscanf(line, "%d %d %d", &r.destStart, &r.sourceStart, &r.size)
	return r
}

// /!\ only works for sorted []rang3
func computeDest(corrMap []rang3, source int) int {
	loc := source
	for _, r := range corrMap {
		if (source < r.sourceStart) || (source > (r.sourceStart + r.size - 1)) {
			continue
		}

		loc = r.destStart - r.sourceStart + source
		break
	}

	return loc
}

func (a almanach) sortMappings() {
	slices.SortFunc(a.seedsToSoil, sortRange)
	slices.SortFunc(a.soilToFertilizer, sortRange)
	slices.SortFunc(a.fertilizerToWater, sortRange)
	slices.SortFunc(a.waterToLight, sortRange)
	slices.SortFunc(a.lightToTemperature, sortRange)
	slices.SortFunc(a.temperatureToHumidity, sortRange)
	slices.SortFunc(a.humidityToLocation, sortRange)
}

func (d *day) loadData() almanach {
	input := utils.LoadInput(d.inputFile)
	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)

	almanach := almanach{
		seeds:                 make([]string, 0),
		seedsToSoil:           make([]rang3, 0),
		soilToFertilizer:      make([]rang3, 0),
		fertilizerToWater:     make([]rang3, 0),
		waterToLight:          make([]rang3, 0),
		lightToTemperature:    make([]rang3, 0),
		temperatureToHumidity: make([]rang3, 0),
		humidityToLocation:    make([]rang3, 0),
	}
	mapsToUse := make([]*[]rang3, 0)
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
		*mapsToUse[ind-1] = append(*mapsToUse[ind-1], rangeFromString(line))
	}
	almanach.sortMappings()
	return almanach
}
