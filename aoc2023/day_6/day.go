package day_6

import (
	"advent-of-code/aoc"
	"advent-of-code/aoc/utils"
	"bufio"
	"log"
	"math"
	"strconv"
	"strings"
)

type day struct {
	inputFile string
}

var _ aoc.EventRunner = (*day)(nil)

func Create() aoc.EventRunner {
	return create("aoc2023/day_6/input.txt")
}

func create(inputFile string) *day {
	return &day{
		inputFile: inputFile,
	}
}

func (d *day) Execute() {
	log.Println("Day 6")
	log.Println("Q1:", d.Question1())
	log.Println("Q2:", d.Question2())
}

func (d *day) Question1() string {
	m, _ := d.loadData()
	tot := 0
	for t, r := range m {
		// compute nb of possible wins
		//var possible int
		//tot *= possible
		totValues := (t + 1) / 2
		bump := false
		if math.Mod(float64(t+1), 2) != 0 {
			totValues++
			bump = true
		}
		possible := 0
		for i := 0; i < totValues; i++ {
			if computeDistForTime(i, t) > r {
				possible++
			}
		}

		possible *= 2
		if bump {
			possible--
		}
		if tot == 0 {
			tot = possible
		} else {
			tot *= possible
		}
	}

	return strconv.Itoa(tot)
}

func computeDistForTime(time int, limit int) int {
	diff := 0
	if limit-time > 0 {
		diff = limit - time
	}
	return time * diff
}

func (d *day) Question2() string {
	_, v := d.loadData()
	// compute nb of possible wins
	//var possible int
	//tot *= possible
	totValues := (v[0] + 1) / 2
	bump := false
	if math.Mod(float64(v[0]+1), 2) != 0 {
		totValues++
		bump = true
	}

	possible := 0
	for i := 0; i < totValues; i++ {
		if computeDistForTime(i, v[0]) > v[1] {
			possible++
		}
	}

	possible *= 2
	if bump {
		possible--
	}
	return strconv.Itoa(possible)
}

func (d *day) loadData() (map[int]int, []int) {
	input := utils.LoadInput(d.inputFile)
	fileScanner := bufio.NewScanner(input)

	fileScanner.Split(bufio.ScanLines)

	timeDistMap := make(map[int]int)
	data := make([][]int, 0)
	res := make([]int, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()

		line = strings.Join(strings.Fields(line), " ")

		base := strings.Split(line, ":")
		tmp := make([]int, 0)

		intVal, _ := strconv.Atoi(strings.Join(strings.Fields(strings.TrimSpace(base[1])), ""))
		res = append(res, intVal)
		for _, s := range strings.Split(strings.TrimSpace(base[1]), " ") {
			v, _ := strconv.Atoi(s)
			tmp = append(tmp, v)
		}
		data = append(data, tmp)
	}

	for i, time := range data[0] {
		timeDistMap[time] = data[1][i]
	}

	return timeDistMap, res
}
