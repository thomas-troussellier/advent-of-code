package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Day13() {
	fmt.Println("Day 13")

	// load exercise data
	seatingPreferences := loadPreferences("aoc2015/inputs/day_13.log")

	perAttendee := arrangeToMap(seatingPreferences)

	fmt.Println("Q1. What is the total change in happiness for the optimal seating arrangement of the actual guest list?")
	sum := computeBestHappiness(perAttendee)
	fmt.Printf("A1. best seating happiness: %d\n", sum)

	// add yourself with no happiness overall
	yourHappiness := make(map[string]int)
	for name := range perAttendee {
		perAttendee[name]["You"] = 0
		yourHappiness[name] = 0
	}
	perAttendee["You"] = yourHappiness

	// recalcule arrangements with you included
	fmt.Println("Q2. What is the total change in happiness for the optimal seating arrangement that actually includes yourself?")
	newSum := computeBestHappiness(perAttendee)
	fmt.Printf("A2. difference from previous best: %d\n", newSum)
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

		//fmt.Printf("seating: %v, hap: %d\n", seating, tempHappiness)
	}

	//fmt.Printf("best seating: %v, hap: %d\n", bSeating, *bestHappiness)

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
		//fmt.Printf("\tseating: %v, hap: %d\n", tempS, tempHap)
	}

	return bSeating, *bestHappiness
}

func bestNextTo(attendees map[string]map[string]int, seated []string, nextTo string, currentHappiness int) ([]string, int) {
	possibleSeats := make([]string, 0)
	previousAtt := seated[len(seated)-1]
	seated = append(seated, nextTo)
	happ := currentHappiness + attendees[previousAtt][nextTo] + attendees[nextTo][previousAtt]
	for name := range attendees[nextTo] {
		if contains(seated, name) {
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
	readFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
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
