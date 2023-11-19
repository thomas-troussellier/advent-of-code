package day_14

import (
	"advent-of-code/aoc"
	"advent-of-code/aoc/utils"
	"bufio"
	"fmt"
	"log"
	"strconv"
)

type day struct {
	inputFile string
}

var _ aoc.EventRunner = (*day)(nil)

func Create() aoc.EventRunner {
	return create("aoc2015/day_14/input.txt")
}

func create(inputFile string) *day {
	return &day{
		inputFile: inputFile,
	}
}

func (d *day) Execute() {
	log.Println("Day 14")
	log.Println("Q1:", d.Question1())
	log.Println("Q2:", d.Question2())
}

func (d *day) Question1() string {
	reindeersData := loadReindeersData(d.inputFile)

	return strconv.Itoa(race(reindeersData, 2503))
}

func (d *day) Question2() string {
	reindeersData := loadReindeersData(d.inputFile)
	standings := raceWithPoints(reindeersData, 2503)
	_, points := computeWinner(standings)

	return strconv.Itoa(points)
}

func computeWinner(standings map[string]int) (winner string, points int) {
	for name := range standings {
		if standings[name] > points {
			points = standings[name]
			winner = name
		}
	}

	return
}

func raceWithPoints(reindeersData []*reindeer, runDuration int) map[string]int {
	resultsTimeline := make(map[string]map[int]int)
	standings := make(map[string]int)

	for _, r := range reindeersData {
		resultsTimeline[r.name] = runFor(r, runDuration)
	}

	for i := 0; i < runDuration; i++ {
		maxDistance := 0
		maxNames := make([]string, 0)
		for name, timeline := range resultsTimeline {
			if timeline[i] > maxDistance {
				maxDistance = timeline[i]
				maxNames = []string{name}
			} else if timeline[i] == maxDistance {
				maxNames = append(maxNames, name)
			}
		}

		for _, name := range maxNames {
			standings[name] += 1
		}
	}

	return standings
}

func race(reindeersData []*reindeer, runDuration int) int {
	maxDistance := 0

	for _, r := range reindeersData {
		timeline := runFor(r, runDuration)

		if timeline[runDuration-1] > maxDistance {
			maxDistance = timeline[runDuration-1]
		}
	}

	return maxDistance
}

func runFor(r *reindeer, runDuration int) map[int]int {
	timeline := make(map[int]int)
	elapsed := 0
	distance := 0

	currentRun := 0
	currentRest := 0
	stateRun := true

	for elapsed < runDuration {
		if stateRun {
			currentRun += 1
			distance += r.fly

			if currentRun == r.flyDuration {
				stateRun = false
				currentRun = 0
			}
		} else {
			currentRest += 1
			if currentRest == r.restDuration {
				stateRun = true
				currentRest = 0
			}
		}
		timeline[elapsed] = distance
		elapsed++
	}

	return timeline
}

type reindeer struct {
	name         string
	fly          int
	flyDuration  int
	restDuration int
}

func loadReindeersData(fileName string) []*reindeer {
	readFile := utils.LoadInput(fileName)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	reindeers := make([]*reindeer, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		var name string
		var fly, flyDuration, restDuration int

		fmt.Sscanf(line, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &name, &fly, &flyDuration, &restDuration)

		reindeers = append(reindeers, &reindeer{
			name:         name,
			fly:          fly,
			flyDuration:  flyDuration,
			restDuration: restDuration,
		})
	}

	readFile.Close()

	return reindeers
}
