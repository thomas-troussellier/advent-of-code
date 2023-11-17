package days

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

func Day15() {
	fmt.Println("Day 15")
	_, occMap := loadSensorFormation("aoc2022/inputs/day_15.log")
	fmt.Println("Q1. In the row where y=2000000, how many positions cannot contain a beacon?", occMap[2000000])
}

func countNotBeacons(floorMap map[int]map[int]rune, line int) (notBeacon int) {
	for _, v := range floorMap[line] {
		if v != 'B' {
			notBeacon++
		}
	}
	return notBeacon
}

func printFloor(floorMap map[int]map[int]string) {
	kY := make([]int, 0)
	minX, maxX := -2, 25
	for key := range floorMap {
		kY = append(kY, key)
	}
	sort.Ints(kY)

	for _, y := range kY {
		kX := make([]int, 0)
		for key := range floorMap[y] {
			kX = append(kX, key)
		}
		sort.Ints(kX)
		fmt.Printf("%3d ", y)
		if len(kX) == 0 {
			fmt.Println()
			continue
		}

		for x := minX; x <= maxX; x++ {
			if v, ok := floorMap[y][x]; !ok {
				fmt.Print(" ")
			} else {
				fmt.Print(v)
			}
		}

		fmt.Println()
	}
}

func loadSensorFormation(fileName string) (map[int]map[int]rune, map[int]int) {
	readFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	floorMap := make(map[int]map[int]rune)
	occMap := make(map[int]int)

	for fileScanner.Scan() {
		currentLine := fileScanner.Text()

		var sX, sY, bX, bY int
		_, err := fmt.Sscanf(currentLine, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sX, &sY, &bX, &bY)
		if err != nil {
			fmt.Println(err)
		}

		setObj(floorMap, 'S', sX, sY)
		occMap[sY] = occMap[sY] + 1
		setObj(floorMap, 'B', bX, bY)

		dist := taxicabDist(sX, sY, bX, bY)
		for y := dist; y >= 0; y-- {
			for x := 0; x <= dist-y; x++ {
				occMap[sY+y] = occMap[sY+y] + 2
				//setObjIfEmpty(floorMap, '.', sX+x, sY+y)
				//setObjIfEmpty(floorMap, '.', sX-x, sY+y)
				occMap[sY-y] = occMap[sY-y] + 2
				//setObjIfEmpty(floorMap, '.', sX-x, sY-y)
				//setObjIfEmpty(floorMap, '.', sX+x, sY-y)
			}
		}

	}

	readFile.Close()
	return floorMap, occMap
}

func taxicabDist(sX, sY, bX, bY int) int {
	return int(math.Abs(float64(sX)-float64(bX)) + math.Abs(float64(sY)-float64(bY)))
}

func setObj(floorMap map[int]map[int]rune, kind rune, x, y int) {
	if _, ok := floorMap[y]; !ok {
		floorMap[y] = make(map[int]rune)
	}
	floorMap[y][x] = kind
}

func setObjIfEmpty(floorMap map[int]map[int]rune, kind rune, x, y int) {
	if _, ok := floorMap[y]; !ok {
		floorMap[y] = make(map[int]rune)
	}

	if _, ok := floorMap[y][x]; !ok {
		floorMap[y][x] = kind
	}
}
