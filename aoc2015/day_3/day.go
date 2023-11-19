package day_3

import (
	"advent-of-code/aoc"
	"advent-of-code/aoc/utils"
	"bufio"
	"log"
	"os"
	"strconv"
)

type day struct {
	inputFile string
}

var _ aoc.EventRunner = (*day)(nil)

func Create() aoc.EventRunner {
	return create("aoc2015/day_3/input.txt")
}

func create(inputFile string) *day {
	return &day{
		inputFile: inputFile,
	}
}

func (d *day) Execute() {
	log.Println("Day 3")
	log.Println("Q1:", d.Question1())
	log.Println("Q2:", d.Question2())
}

func (d *day) Question1() string {
	homes := loadOneSanta(d.inputFile)

	return strconv.Itoa(len(homes))
}

func (d *day) Question2() string {
	homes := loadTwoSanta(d.inputFile)

	return strconv.Itoa(len(homes))
}

type house struct {
	x        int
	y        int
	presents int
}

func (h *house) GetDirectionHouse(dir string, directions map[int]map[int]*house) (toHouse *house, found bool) {

	x, y := h.GetCoordinatesForDir(dir)

	var (
		tmp map[int]*house
		ok  bool
	)
	if tmp, ok = directions[x]; !ok {
		return &house{x: x, y: y}, false
	}

	if _, ok = tmp[y]; !ok {
		return &house{x: x, y: y}, false
	}
	return tmp[y], true
}

func (h *house) GetCoordinatesForDir(dir string) (x int, y int) {
	switch dir {
	case "^":
		x = h.x
		y = h.y + 1
	case ">":
		x = h.x + 1
		y = h.y
	case "v":
		x = h.x
		y = h.y - 1
	case "<":
		x = h.x - 1
		y = h.y
	}
	return
}

type houses []*house

func loadOneSanta(fileName string) houses {
	readFile := utils.LoadInput(fileName)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanRunes)

	houseList := make(houses, 0)
	directions := make(map[int]map[int]*house)

	// starting point, should get a present
	fromHouse := &house{presents: 1, x: 0, y: 0}
	houseList = append(houseList, fromHouse)
	directions[fromHouse.x] = make(map[int]*house)
	directions[fromHouse.x][fromHouse.y] = fromHouse

	temp := fromHouse
	for fileScanner.Scan() {
		currentDirection := fileScanner.Text()

		// get the next house
		toHouse, found := temp.GetDirectionHouse(currentDirection, directions)

		// if it was not visited yet
		if !found {
			if _, ok := directions[toHouse.x]; !ok {
				directions[toHouse.x] = make(map[int]*house)
			}

			directions[toHouse.x][toHouse.y] = toHouse

			// add to list
			houseList = append(houseList, toHouse)
		}

		toHouse.presents += 1

		temp = toHouse
	}

	readFile.Close()

	return houseList
}

func loadTwoSanta(fileName string) houses {
	readFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanRunes)

	houseList := make(houses, 0)
	directions := make(map[int]map[int]*house)

	// starting point, should get a present
	fromHouse := &house{presents: 1, x: 0, y: 0}
	houseList = append(houseList, fromHouse)
	directions[fromHouse.x] = make(map[int]*house)
	directions[fromHouse.x][fromHouse.y] = fromHouse

	tempMan := fromHouse
	tempBot := fromHouse
	toMove := santa_man
	for fileScanner.Scan() {
		currentDirection := fileScanner.Text()
		var temp *house

		switch toMove {
		case santa_man:
			temp = tempMan
		case santa_robot:
			temp = tempBot
		}
		// get the next house
		toHouse, found := temp.GetDirectionHouse(currentDirection, directions)

		// if it was not visited yet
		if !found {
			if _, ok := directions[toHouse.x]; !ok {
				directions[toHouse.x] = make(map[int]*house)
			}

			directions[toHouse.x][toHouse.y] = toHouse

			// add to list
			houseList = append(houseList, toHouse)
		}

		toHouse.presents += 1

		switch toMove {
		case santa_man:
			tempMan = toHouse
			toMove = santa_robot
		case santa_robot:
			tempBot = toHouse
			toMove = santa_man
		}
	}

	readFile.Close()

	return houseList
}

const (
	santa_man   = "SANTA"
	santa_robot = "SANTA_BOT"
)
