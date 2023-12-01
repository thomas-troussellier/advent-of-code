package day_15

import (
	"advent-of-code/aoc"
	"advent-of-code/aoc/utils"
	"bufio"
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
)

type day struct {
	inputFile string
}

var _ aoc.EventRunner = (*day)(nil)

func Create() aoc.EventRunner {
	return create("aoc2022/day_15/input.txt")
}

func create(inputFile string) *day {
	return &day{
		inputFile: inputFile,
	}
}

func (d *day) Execute() {
	log.Println("Day 15")
	log.Println("Q1:", d.Question1())
	log.Println("Q2:", d.Question2())
}

func (d *day) Question1() string {
	s, b := d.loadData()
	return strconv.Itoa(computeNonAvailableSpaces(s, b, 2000000))
}

func (d *day) Question2() string {

	return "Question2"
}

func taxicabDist(sX, sY, bX, bY int) int {
	return int(math.Abs(float64(sX)-float64(bX)) + math.Abs(float64(sY)-float64(bY)))
}

type sensor struct {
	coord       coordinates
	taxicabdist int
}

type coordinates struct {
	x, y int
}

type radius struct {
	from, to int
}

var outsideRange error

// returns the x range for positions in the manhattan distance for the given y value or error if outside range
func (s sensor) manhattanDistPosNb(yCoord int) (rad radius, err error) {
	// compute if y is in sensor range
	diffVal := s.coord.y - yCoord
	absValue := math.Abs(float64(diffVal))
	if absValue > float64(s.taxicabdist) {
		return rad, outsideRange
	}
	return radius{from: s.coord.x - (s.taxicabdist - int(absValue)), to: s.coord.x + (s.taxicabdist - int(absValue))}, nil
}

func computeNonAvailableSpaces(sensors []sensor, beacons []coordinates, yCoord int) int {
	pos := make([]radius, 0)
	for _, sen := range sensors {
		if r, err := sen.manhattanDistPosNb(yCoord); err == nil {
			pos = append(pos, r)
		}
	}

	// then consolidate ranges
	slices.SortFunc(pos, func(a, b radius) int {
		return a.from - b.from
	})

	if len(pos) == 0 {
		return 0
	}
	// regroup possible ranges
	conso := make(map[int]any, 0)
	for i := range pos {
		for j := pos[i].from; j <= pos[i].to; j++ {
			conso[j] = struct{}{}
		}
	}
	for _, beacon := range beacons {
		if beacon.y == yCoord {
			delete(conso, beacon.x)
		}
	}

	return len(conso)
}

func (d *day) loadData() ([]sensor, []coordinates) {
	readFile := utils.LoadInput(d.inputFile)

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	knownSensors := make([]sensor, 0)
	knownBeacons := make([]coordinates, 0)
	for fileScanner.Scan() {
		currentLine := fileScanner.Text()

		var sen sensor
		var beacon coordinates
		_, err := fmt.Sscanf(currentLine, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sen.coord.x, &sen.coord.y, &beacon.x, &beacon.y)
		if err != nil {
			fmt.Println(err)
		}

		sen.taxicabdist = taxicabDist(sen.coord.x, sen.coord.y, beacon.x, beacon.y)
		knownSensors = append(knownSensors, sen)
		knownBeacons = append(knownBeacons, beacon)
	}

	return knownSensors, knownBeacons
}
