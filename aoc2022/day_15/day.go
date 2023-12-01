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
	lineToRead := 2000000
	return strconv.Itoa(nbUnavailableSpacesOnLine(s, b, lineToRead))
}

func nbUnavailableSpacesOnLine(sensors []sensor, beacons []coordinates, yCoord int) int {
	nonAv := computeNonAvailableSpaces(sensors, beacons, yCoord)
	// we need to remove beacon spaces
	bList := make([]int, 0)
	for _, beac := range beacons {
		if beac.y != yCoord {
			continue
		}
		if !slices.Contains(bList, beac.x) {
			bList = append(bList, beac.x)
		}
	}
	slices.Sort(bList)
	sum := 0
	for _, na := range nonAv {
		sum += (na.to - na.from)
		if na.from < 0 && na.to > 0 {
			sum++
		}
		for _, v := range bList {
			if v >= na.from && v <= na.to {
				sum--
			}
		}
	}
	return sum
}

func (d *day) Question2() string {
	s, b := d.loadData()

	return strconv.Itoa(findFrequency(s, b, 4000000))
}

func findFrequency(s []sensor, b []coordinates, yIt int) int {
	for y := 0; y <= yIt; y++ {
		available := computeAvailableSpaces(s, b, y)
		if len(available) > 0 {
			return (4000000 * available[0].from) + y
		}
	}
	log.Fatal("could not find any avaliable space")
	return -1
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

func (r radius) String() string {
	return fmt.Sprintf("{f:%d,t:%d}", r.from, r.to)
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
	return radius{from: (s.coord.x - (s.taxicabdist - int(absValue))), to: (s.coord.x + (s.taxicabdist - int(absValue)))}, nil
}

func computeNonAvailableSpaces(sensors []sensor, beacons []coordinates, yCoord int) []radius {
	pos := make([]radius, 0)
	for _, sen := range sensors {
		if r, err := sen.manhattanDistPosNb(yCoord); err == nil {
			pos = append(pos, r)
		}
	}
	if len(pos) == 0 {
		return pos
	}

	// then consolidate ranges
	slices.SortFunc(pos, func(a, b radius) int {
		return a.from - b.from
	})

	// regroup possible ranges
	consolided := make([]radius, 0)
	consolided = append(consolided, pos[0])
	for _, v := range pos[1:] {
		if v.from <= consolided[len(consolided)-1].to {
			if v.to < consolided[len(consolided)-1].to {
				continue
			} else {
				consolided[len(consolided)-1].to = v.to
			}
		} else {
			consolided = append(consolided, v)
		}
	}
	return consolided
}

func computeAvailableSpaces(sensors []sensor, beacons []coordinates, yCoord int) []radius {
	nonAvailable := computeNonAvailableSpaces(sensors, beacons, yCoord)

	if len(nonAvailable) == 1 {
		return []radius{}
	}

	avail := make([]radius, 0)
	for i := range nonAvailable[:len(nonAvailable)-1] {
		avail = append(avail, radius{from: nonAvailable[i].to + 1, to: nonAvailable[i+1].from - 1})
	}

	return avail
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

	slices.SortFunc(knownSensors, func(a, b sensor) int {
		return (a.coord.y + a.taxicabdist) - (b.coord.x + a.taxicabdist)
	})

	slices.SortFunc(knownBeacons, func(a, b coordinates) int {
		return a.x - b.x
	})

	return knownSensors, knownBeacons
}
