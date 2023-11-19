package day_13

import (
	"advent-of-code/aoc"
	"advent-of-code/aoc/utils"
	"bufio"
	"fmt"
	"log"
	"slices"
	"strconv"
)

type day struct {
	inputFile string
}

var _ aoc.EventRunner = (*day)(nil)

func Create() aoc.EventRunner {
	return create("aoc2015/day_13/input.txt")
}

func create(inputFile string) *day {
	return &day{
		inputFile: inputFile,
	}
}

func (d *day) Execute() {
	log.Println("Day 13")
	log.Println("Q1:", d.Question1())
	log.Println("Q2:", d.Question2())
}

func (d *day) Question1() string {
	seatingPreferences := loadPreferences(d.inputFile)

	perAttendee := arrangeToMap(seatingPreferences)

	return strconv.Itoa(computeBestHappiness(perAttendee))
}

func (d *day) Question2() string {
	seatingPreferences := loadPreferences(d.inputFile)

	perAttendee := arrangeToMap(seatingPreferences)
	yourHappiness := make(map[string]int)
	for name := range perAttendee {
		perAttendee[name]["You"] = 0
		yourHappiness[name] = 0
	}
	perAttendee["You"] = yourHappiness

	return strconv.Itoa(computeBestHappiness(perAttendee))
}

func computeBestHappiness(perAttendee map[string]map[string]int) int {
	bestHappiness := new(int)
	var bSeating []string

	// for each attendee, compute best seating arrangement
	for attendee := range perAttendee {
		// keep best happiness
		seating, tempHappiness := bestSeating(perAttendee, attendee)

		if bestHappiness == nil || ((tempHappiness > *bestHappiness) && len(seating) >= len(bSeating)) {
			*bestHappiness = tempHappiness
			bSeating = seating
		}
	}

	return *bestHappiness
}

func bestSeating(perAttendee map[string]map[string]int, attendee string) ([]string, int) {
	bestHappiness := new(int)
	bSeating := make([]string, 0)

	for name := range perAttendee[attendee] {
		currentHap := 0
		tempS, tempHap := bestNextTo(perAttendee, []string{attendee}, name, currentHap)

		if bestHappiness == nil || ((tempHap > *bestHappiness) && len(tempS) >= len(bSeating)) {
			*bestHappiness = tempHap
			bSeating = tempS
		}
	}

	return bSeating, *bestHappiness
}

func bestNextTo(attendees map[string]map[string]int, seated []string, nextTo string, currentHappiness int) ([]string, int) {
	possibleSeats := make([]string, 0)
	previousAtt := seated[len(seated)-1]
	seated = append(seated, nextTo)
	happ := currentHappiness + attendees[previousAtt][nextTo] + attendees[nextTo][previousAtt]
	for name := range attendees[nextTo] {
		if slices.Contains(seated, name) {
			continue
		}
		possibleSeats = append(possibleSeats, name)
	}

	if len(possibleSeats) == 0 {
		return seated, happ + attendees[nextTo][seated[0]] + attendees[seated[0]][nextTo]
	}

	bestHappiness := new(int)
	bSeating := seated

	for _, name := range possibleSeats {
		tempS, tempHap := bestNextTo(attendees, seated, name, happ)

		if bestHappiness == nil || ((tempHap > *bestHappiness) && len(tempS) >= len(bSeating)) {
			*bestHappiness = tempHap
			bSeating = tempS
		}
	}

	return bSeating, *bestHappiness
}

func arrangeToMap(prefs []string) map[string]map[string]int {
	perAttendee := make(map[string]map[string]int)

	for _, attendeePref := range prefs {
		var att1, att2, mood string
		var hap int

		fmt.Sscanf(attendeePref, "%s would %s %d happiness units by sitting next to %s.", &att1, &mood, &hap, &att2)
		// hack cause sscanf %s gets space separated tokens and the '.' int the format gets sent in the last %s
		att2 = att2[:len(att2)-1]

		switch mood {
		case "lose":
			hap = hap * -1
		}

		if _, ok := perAttendee[att1]; !ok {
			perAttendee[att1] = make(map[string]int)
			perAttendee[att1][att2] = hap
		} else {
			perAttendee[att1][att2] = hap
		}
	}

	return perAttendee
}

func loadPreferences(fileName string) []string {
	readFile := utils.LoadInput(fileName)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	prefs := make([]string, 0)
	for fileScanner.Scan() {
		pref := fileScanner.Text()
		prefs = append(prefs, pref)
	}

	readFile.Close()

	return prefs
}
