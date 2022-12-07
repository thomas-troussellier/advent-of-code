package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Day14() {
	fmt.Println("Day 14")

	// load exercise data
	reindeersData := loadReindeersData("aoc2015/inputs/day_14.log")

	fmt.Println("Q1. After exactly 2503 seconds, what distance has the winning reindeer traveled?")
	maxDistance := race(reindeersData, 2503)
	fmt.Printf("A1. Furthest distance traveled: %d\n", maxDistance)

	fmt.Println("Q2. After exactly 2503 seconds, how many points does the winning reindeer have?")
	standings := raceWithPoints(reindeersData, 2503)
	winner, points := computeWinner(standings)
	fmt.Printf("A2. Highest score: %d by %s\n", points, winner)
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

	fmt.Printf("standings: %v\n", standings)

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
	readFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
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
