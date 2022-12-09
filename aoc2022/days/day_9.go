package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Day9() {
	fmt.Println("Day 9")

	directions := loadHeadDirections("aoc2022/inputs/day_9.log")

	fmt.Println("Q1. How many positions does the tail of the rope visit at least once?")
	tailVisitedOnce := computeRopeMovment(directions, 2)
	fmt.Printf("A1. Nb of positions visited at least once by tail: %d\n", len(tailVisitedOnce))

	fmt.Println("Q2. Simulate your complete series of motions on a larger rope with ten knots. How many positions does the tail of the rope visit at least once?")
	tailVisitedOnce = computeRopeMovment(directions, 10)
	fmt.Printf("A2. Nb of positions visited at least once by tail: %d\n", len(tailVisitedOnce))
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
