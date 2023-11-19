package day_9

import (
	"advent-of-code/aoc"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type day struct {
	inputFile string
}

var _ aoc.EventRunner = (*day)(nil)

func Create() aoc.EventRunner {
	return create("aoc2022/day_9/input.txt")
}

func create(inputFile string) *day {
	return &day{
		inputFile: inputFile,
	}
}

func (d *day) Execute() {
	log.Println("Day 9")
	log.Println("Q1:", d.Question1())
	log.Println("Q2:", d.Question2())
}

func (d *day) Question1() string {
	directions := loadHeadDirections(d.inputFile)

	return strconv.Itoa(len(computeRopeMovment(directions, 2)))
}

func (d *day) Question2() string {
	directions := loadHeadDirections(d.inputFile)

	return strconv.Itoa(len(computeRopeMovment(directions, 10)))
}

func computeRopeMovment(directions []*dir, ropeSize int) map[string]coordinates {
	rope := make(map[int]coordinates)
	for i := 0; i < ropeSize; i++ {
		rope[i] = coordinates{}
	}

	tailVisited := map[string]coordinates{rope[ropeSize-1].String(): rope[ropeSize-1]}

	for _, d := range directions {
		d := d
		for i := 0; i < d.distance; i++ {
			for index := 0; index < ropeSize; index++ {
				var knot = rope[index]
				if index == 0 {
					switch d.direction {
					case "U":
						knot.y = rope[index].y + 1
					case "D":
						knot.y = rope[index].y - 1
					case "L":
						knot.x = rope[index].x - 1
					case "R":
						knot.x = rope[index].x + 1
					}
					rope[index] = knot
					continue
				}

				knot = computeNewTailPos(rope[index], rope[index-1])

				if index == ropeSize-1 {
					s := knot.String()
					if _, ok := tailVisited[s]; !ok {
						tailVisited[s] = knot
					}
				}

				rope[index] = knot
			}
		}
	}

	return tailVisited
}

func computeNewTailPos(tailPosition coordinates, headPosition coordinates) coordinates {
	if adjacent(tailPosition, headPosition) {
		return tailPosition
	}
	newTailPos := tailPosition

	if headPosition.x > tailPosition.x {
		newTailPos.x += 1
	} else if headPosition.x < tailPosition.x {
		newTailPos.x -= 1
	}

	if headPosition.y > tailPosition.y {
		newTailPos.y += 1
	} else if headPosition.y < tailPosition.y {
		newTailPos.y -= 1
	}

	return newTailPos
}

func adjacent(tailPosition, headPosition coordinates) bool {
	movesToHead := coordinates{
		x: (headPosition.x - tailPosition.x),
		y: (headPosition.y - tailPosition.y),
	}

	if (movesToHead.x < -1 || movesToHead.x > 1) || (movesToHead.y < -1 || movesToHead.y > 1) {
		return false
	}

	return true
}

type dir struct {
	direction string
	distance  int
}

func (d dir) String() string {
	return fmt.Sprintf("%s,%d", d.direction, d.distance)
}

type coordinates struct {
	x int
	y int
}

func (c coordinates) String() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}

func loadHeadDirections(fileName string) []*dir {
	readFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	dirs := make([]*dir, 0)

	for fileScanner.Scan() {
		currentLine := fileScanner.Text()

		var (
			direction string
			distance  int
		)

		fmt.Sscanf(currentLine, "%s %d", &direction, &distance)

		dirs = append(dirs, &dir{
			direction: direction,
			distance:  distance,
		})
	}

	readFile.Close()

	return dirs
}
