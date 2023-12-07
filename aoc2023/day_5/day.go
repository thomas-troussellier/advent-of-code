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
			res <- computeLocForSeedRange(al, start, size)
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

func computeLocForSeed(al *almanach, seed int) int {
	mapping := [][]rang3{al.seedsToSoil, al.soilToFertilizer, al.fertilizerToWater, al.waterToLight, al.lightToTemperature, al.temperatureToHumidity, al.humidityToLocation}
	for i := range mapping {
		seed = computeDest(mapping[i], seed)
	}

	return seed
}

func computeLocForSeedRange(al *almanach, seedStart, size int) int {
	min := 0

	for s := seedStart; s < seedStart+size; s++ {
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
	deltas      []int
	sourceStart int
	delta       int
	size        int
}

func (r rang3) String() string {
	//return fmt.Sprintf("%d - %d => %d - %d", r.sourceStart, r.sourceStart+r.size-1, r.sourceStart+r.delta, r.sourceStart+r.size-1+r.delta)
	return fmt.Sprintf("(%d - %d {%v} %d)", r.sourceStart, r.sourceStart+r.size-1, r.deltas, r.delta)
}

var sortSourceRange = func(a rang3, b rang3) int {
	return a.sourceStart - b.sourceStart
}

func rangeFromString(line string) rang3 {
	r := rang3{}
	fmt.Sscanf(line, "%d %d %d", &r.delta, &r.sourceStart, &r.size)
	r.delta = r.delta - r.sourceStart
	r.deltas = append(r.deltas, r.delta)
	return r
}

// /!\ only works for sorted []rang3
func computeDest(corrMap []rang3, source int) int {
	loc := source
	for _, r := range corrMap {
		if source < r.sourceStart {
			break
		}

		if source > (r.sourceStart + r.size - 1) {
			continue
		}

		loc = source + r.delta
		break
	}

	return loc
}

func (a *almanach) sortMappings() {
	slices.SortFunc(a.seedsToSoil, sortSourceRange)
	slices.SortFunc(a.soilToFertilizer, sortSourceRange)
	slices.SortFunc(a.fertilizerToWater, sortSourceRange)
	slices.SortFunc(a.waterToLight, sortSourceRange)
	slices.SortFunc(a.lightToTemperature, sortSourceRange)
	slices.SortFunc(a.temperatureToHumidity, sortSourceRange)
	slices.SortFunc(a.humidityToLocation, sortSourceRange)
}

func (d *day) loadData() *almanach {
	input := utils.LoadInput(d.inputFile)
	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)

	almanach := &almanach{
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

func consolidate(base []rang3, in []rang3) []rang3 {
	//log.Println("base", base)
	//log.Println("in", in)

	consolided := make([]rang3, len(base))
	copy(consolided, base)

	for _, new := range in {
		consolided = consolidateList(consolided, &new, len(base[0].deltas))
	}

	for i, v := range consolided {
		if len(v.deltas) < len(base[0].deltas)+1 {
			consolided[i].deltas = append(consolided[i].deltas, 0)
		}
	}

	slices.SortFunc(consolided, sortSourceRange)
	//log.Println("consolidate", consolided)
	return consolided
}

func consolidateList(toModify []rang3, new *rang3, base int) []rang3 {
	if new == nil {
		return toModify
	}
	consolided := make([]rang3, len(toModify))
	copy(consolided, toModify)

	for i, ra := range toModify {
		if new.sourceStart > (ra.sourceStart + ra.size - 1) {
			continue
		}
		newRanges, toCheck := consolidateOne(ra, *new, base)
		// replace ra index with newRanges
		if i < len(consolided)-1 {
			consolided = append(consolided[:i], append(newRanges, consolided[i+1:]...)...)
		} else {
			consolided = append(consolided[:i], newRanges...)
		}
		if toCheck == nil {
			break
		}
		consolided = consolidateList(consolided, toCheck, base)
		break
	}

	slices.SortFunc(consolided, sortSourceRange)

	return consolided
}

func consolidateOne(current rang3, new rang3, base int) ([]rang3, *rang3) {
	consolided := make([]rang3, 0)

	currentStart := current.sourceStart
	currentEnd := current.sourceStart + current.size - 1

	newStart := new.sourceStart
	newEnd := new.sourceStart + new.size - 1

	var toCheck *rang3

	//log.Printf("cS: %d, nS: %d, cE:%d, nE:%d, deltas: cD:%d, nD:%d\n", currentStart, newStart, currentEnd, newEnd, current.deltas, new.deltas)
	if newStart < currentStart {
		//log.Println("newStart < currentStart")
		if newEnd < currentStart {
			//log.Println("newEnd < currentStart")
			consolided = append(consolided, rang3{sourceStart: newStart, size: new.size, delta: new.delta, deltas: append(make([]int, base), new.delta)})
			consolided = append(consolided, rang3{sourceStart: currentStart, size: current.size, delta: current.delta, deltas: current.deltas})
		} else if newEnd == currentStart {
			//log.Println("newEnd == currentStart")
			consolided = append(consolided, rang3{sourceStart: newStart, size: new.size - 1, delta: new.delta, deltas: append(make([]int, base), new.delta)})
			consolided = append(consolided, rang3{sourceStart: currentStart, size: 1, delta: new.delta, deltas: append(current.deltas, new.delta)})
			consolided = append(consolided, rang3{sourceStart: currentStart + 1, size: current.size - 1, delta: current.delta, deltas: current.deltas})
		} else {
			//log.Println("newEnd > currentStart")
			if newEnd < currentEnd {
				//log.Println("newEnd < currentEnd")
				consolided = append(consolided, rang3{sourceStart: newStart, size: currentStart - newStart, delta: new.delta, deltas: append(make([]int, base), new.delta)})
				consolided = append(consolided, rang3{sourceStart: currentStart, size: current.size - (currentEnd - newEnd), delta: new.delta, deltas: append(current.deltas, new.delta)})
				consolided = append(consolided, rang3{sourceStart: newEnd + 1, size: (currentEnd - newEnd), delta: current.delta, deltas: current.deltas})
			} else if newEnd == currentEnd {
				//	log.Println("newEnd == currentEnd")
				consolided = append(consolided, rang3{sourceStart: newStart, size: currentStart - newStart, delta: new.delta, deltas: append(make([]int, base), new.delta)})
				consolided = append(consolided, rang3{sourceStart: currentStart, size: current.size, delta: new.delta, deltas: append(current.deltas, new.delta)})
			} else {
				//log.Println("newEnd > currentEnd")
				consolided = append(consolided, rang3{sourceStart: newStart, size: currentStart - newStart, delta: new.delta, deltas: append(make([]int, base), new.delta)})
				consolided = append(consolided, rang3{sourceStart: currentStart, size: current.size, delta: new.delta, deltas: append(current.deltas, new.delta)})
				toCheck = &rang3{sourceStart: currentEnd + 1, size: newEnd - currentEnd, delta: new.delta, deltas: new.deltas}
			}
		}
	} else if newStart == currentStart {
		//log.Println("newStart == currentStart")
		if newEnd < currentEnd {
			//log.Println("newEnd < currentEnd")
			consolided = append(consolided, rang3{sourceStart: newStart, size: new.size, delta: new.delta + current.delta, deltas: append(current.deltas, new.delta)})
			consolided = append(consolided, rang3{sourceStart: newEnd + 1, size: currentEnd - newEnd, delta: current.delta, deltas: current.deltas})
		} else if newEnd == currentEnd {
			//log.Println("newEnd == currentEnd")
			consolided = append(consolided, rang3{sourceStart: currentStart, size: current.size, delta: current.delta + new.delta, deltas: append(current.deltas, new.delta)})
		} else {
			//log.Println("newEnd > currentEnd")
			consolided = append(consolided, rang3{sourceStart: newStart, size: current.size, delta: current.delta + new.delta, deltas: append(current.deltas, new.delta)})
			toCheck = &rang3{sourceStart: currentEnd + 1, size: newEnd - currentEnd, delta: new.delta, deltas: new.deltas}
		}
	} else {
		//log.Println("newStart > currentStart")
		if newStart > currentEnd {
			//log.Println("newStart > currentEnd")
			toCheck = &new
		}
		if newEnd < currentEnd {
			//log.Println("newEnd < currentEnd")
			consolided = append(consolided, rang3{sourceStart: currentStart, size: newStart - currentStart, delta: current.delta, deltas: current.deltas})
			consolided = append(consolided, rang3{sourceStart: newStart, size: new.size, delta: current.delta + new.delta, deltas: append(current.deltas, new.delta)})
			consolided = append(consolided, rang3{sourceStart: newEnd + 1, size: currentEnd - newEnd, delta: current.delta, deltas: current.deltas})
		} else if newEnd == currentEnd {
			//log.Println("newEnd == currentEnd")
			consolided = append(consolided, rang3{sourceStart: currentStart, size: newStart - currentStart, delta: current.delta, deltas: current.deltas})
			consolided = append(consolided, rang3{sourceStart: newStart, size: new.size, delta: current.delta + new.delta, deltas: append(current.deltas, new.delta)})
		} else {
			//log.Println("newEnd > currentEnd")
			consolided = append(consolided, rang3{sourceStart: currentStart, size: newStart - currentStart, delta: current.delta, deltas: current.deltas})
			consolided = append(consolided, rang3{sourceStart: newStart, size: current.size - (newStart - currentStart), delta: new.delta + current.delta, deltas: append(current.deltas, new.delta)})
			toCheck = &rang3{sourceStart: currentEnd + 1, size: newEnd - currentEnd, delta: new.delta, deltas: new.deltas}
		}
	}
	slices.SortFunc(consolided, sortSourceRange)

	//log.Println("consolidateOne", consolided, toCheck)

	return consolided, toCheck
}
