package days

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day14() {

	fmt.Println("Day 14")
	rocks, bottomLimit := loadRockFormation("aoc2022/inputs/day_14.log")

	fmt.Println("Q1. How many units of sand come to rest before sand starts flowing into the abyss below?")
	sandDisposition := makeItFlow(rocks, bottomLimit)
	fmt.Printf("A1. units of sand that came to rest: %d\n", count(sandDisposition))
	//printObstacles(rocks, sandDisposition)

	fmt.Println("Q2. Simulate the falling sand until the source of the sand becomes blocked. How many units of sand come to rest?")
	sandDisposition = makeItFlowWithHardFloor(rocks, bottomLimit+2)
	fmt.Printf("A2. units of sand that came to rest: %d\n", count(sandDisposition))
	//printObstacles(rocks, sandDisposition)
}

func makeItFlow(rocks map[int][]int, bottomLimit int) map[int][]int {
	// sand starts to "flow" at 500,0
	x, y := 500, 0
	sandDisposition := make(map[int][]int)

	// continue until sand starts to go into the abyss
	for y <= bottomLimit {
		// check directly below
		xBis, yBis := down(x, y)
		// if nothing, go there
		if !isThereSmthg(rocks, xBis, yBis) && !isThereSmthg(sandDisposition, xBis, yBis) {
			x = xBis
			y = yBis
			continue
		}
		// if smth,
		// try go left
		xBis, yBis = left(x, y)
		// if nothing, go there
		if !isThereSmthg(rocks, xBis, yBis) && !isThereSmthg(sandDisposition, xBis, yBis) {
			x = xBis
			y = yBis
			continue
		}
		// else go rigth
		xBis, yBis = right(x, y)
		// if nothing, go there
		if !isThereSmthg(rocks, xBis, yBis) && !isThereSmthg(sandDisposition, xBis, yBis) {
			x = xBis
			y = yBis
			continue
		}

		// if we can't go anywhere else, we stay there
		if _, ok := sandDisposition[x]; !ok {
			sandDisposition[x] = make([]int, 0)
		}
		sandDisposition[x] = append(sandDisposition[x], y)

		// and then we produce the next sand unit
		x, y = 500, 0
	}

	return sandDisposition
}

func makeItFlowWithHardFloor(rocks map[int][]int, bottomLimit int) map[int][]int {
	// sand starts to "flow" at 500,0
	x, y := 500, 0
	sandDisposition := make(map[int][]int)

	// continue until sand starts to go into the abyss
	for !contains(sandDisposition[500], 0) {
		// check directly below
		xBis, yBis := down(x, y)

		// hit the floor
		if yBis == bottomLimit {
			// if we can't go anywhere else, we stay there
			if _, ok := sandDisposition[x]; !ok {
				sandDisposition[x] = make([]int, 0)
			}
			sandDisposition[x] = append(sandDisposition[x], y)

			// and then we produce the next sand unit
			x, y = 500, 0
			continue
		}

		// if nothing, go there
		if !isThereSmthg(rocks, xBis, yBis) && !isThereSmthg(sandDisposition, xBis, yBis) {
			x = xBis
			y = yBis
			continue
		}
		// if smth,
		// try go left
		xBis, yBis = left(x, y)
		// if nothing, go there
		if !isThereSmthg(rocks, xBis, yBis) && !isThereSmthg(sandDisposition, xBis, yBis) {
			x = xBis
			y = yBis
			continue
		}
		// else go rigth
		xBis, yBis = right(x, y)
		// if nothing, go there
		if !isThereSmthg(rocks, xBis, yBis) && !isThereSmthg(sandDisposition, xBis, yBis) {
			x = xBis
			y = yBis
			continue
		}

		// if we can't go anywhere else, we stay there
		if _, ok := sandDisposition[x]; !ok {
			sandDisposition[x] = make([]int, 0)
		}
		sandDisposition[x] = append(sandDisposition[x], y)

		// and then we produce the next sand unit
		x, y = 500, 0
	}

	return sandDisposition
}

func isThereSmthg(obstacles map[int][]int, x, y int) bool {
	if _, ok := obstacles[x]; !ok {
		return false
	}
	items := obstacles[x]
	return contains(items, y)
}

func down(x, y int) (int, int) {
	return x, y + 1
}

func left(x, y int) (int, int) {
	return x - 1, y + 1
}

func right(x, y int) (int, int) {
	return x + 1, y + 1
}

func loadRockFormation(fileName string) (map[int][]int, int) {
	readFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	rocks := make(map[int][]int, 0)
	bottomLimit := 0

	for fileScanner.Scan() {
		currentLine := fileScanner.Text()
		// line is x,y -> ...
		coordinates := strings.Split(currentLine, " ")
		var prX, prY int
		for i, part := range coordinates {
			// skip the ->
			if part == "->" {
				continue
			}
			coor := strings.Split(part, ",")
			x, _ := strconv.Atoi(coor[0])
			y, _ := strconv.Atoi(coor[1])

			// bottomLimit is the highest y we get
			if bottomLimit == 0 || y > bottomLimit {
				bottomLimit = y
			}

			if _, ok := rocks[x]; !ok {
				rocks[x] = make([]int, 0)
			}

			rocks[x] = append(rocks[x], y)

			if i != 0 {
				rocks = addInterval(rocks, prX, x, prY, y)
			}
			prX = x
			prY = y
		}
	}

	readFile.Close()

	return rocks, bottomLimit
}

func addInterval(rocks map[int][]int, prX, x, prY, y int) map[int][]int {
	final := rocks

	if prX-x != 0 {
		for i := int(math.Min(float64(prX), float64(x))) + 1; i < int(math.Max(float64(prX), float64(x))); i++ {
			if _, ok := final[i]; !ok {
				final[i] = make([]int, 0)
			}
			final[i] = append(final[i], y)
		}
	} else {
		for i := int(math.Min(float64(prY), float64(y))) + 1; i < int(math.Max(float64(prY), float64(y))); i++ {
			if _, ok := final[x]; !ok {
				final[x] = make([]int, 0)
			}
			final[x] = append(final[x], i)
		}
	}

	return final
}

func contains(list []int, elmt int) bool {
	for _, a := range list {
		if a == elmt {
			return true
		}
	}
	return false
}

func count(sandDisposition map[int][]int) int {
	total := 0

	for _, v := range sandDisposition {
		total += len(v)
	}

	return total
}

func printObstacles(rocks, sand map[int][]int) {
	dim := make(map[int]map[int]string)
	for k, v := range rocks {
		dim[k] = make(map[int]string)
		for _, r := range v {
			dim[k][r] = "#"
		}
	}
	for k, v := range sand {
		if _, ok := dim[k]; !ok {
			dim[k] = make(map[int]string)
		}
		for _, s := range v {
			dim[k][s] = "o"
		}
	}

	yKeys := make([]int, 0)
	for _, v := range dim {
		for k := range v {
			if !contains(yKeys, k) {
				yKeys = append(yKeys, k)
			}
		}
	}
	sort.Ints(yKeys)

	for _, y := range yKeys {
		xKeys := make([]int, 0)
		for x := range dim {
			xKeys = append(xKeys, x)
		}
		sort.Ints(xKeys)

		for i := xKeys[0]; i <= xKeys[len(xKeys)-1]; i++ {
			if o, ok := dim[i][y]; !ok {
				fmt.Print(" ")
			} else {
				fmt.Print(o)
			}
		}
		fmt.Println()
	}
}
