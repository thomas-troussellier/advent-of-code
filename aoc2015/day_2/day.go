package day_2

import (
	"advent-of-code/aoc"
	"advent-of-code/aoc/utils"
	"bufio"
	"log"
	"sort"
	"strconv"
	"strings"
)

type day struct {
	inputFile string
}

var _ aoc.EventRunner = (*day)(nil)

func Create() aoc.EventRunner {
	return create("aoc2015/day_2/input.txt")
}

func create(inputFile string) *day {
	return &day{
		inputFile: inputFile,
	}
}

func (d *day) Execute() {
	log.Println("Day 2")
	log.Println("Q1:", d.Question1())
	log.Println("Q2:", d.Question2())
}

func (d *day) Question1() string {
	presents := loadPresentsData(d.inputFile)
	totalWrapping := 0
	for _, p := range *presents {
		totalWrapping += p.wrappingSurfaceNeeded
	}

	return strconv.Itoa(totalWrapping)
}

func (d *day) Question2() string {
	presents := loadPresentsData(d.inputFile)
	totalRibbon := 0
	for _, p := range *presents {
		totalRibbon += p.ribbonNeeded
	}

	return strconv.Itoa(totalRibbon)
}

type present struct {
	length int
	width  int
	height int

	ribbonNeeded          int
	wrappingSurfaceNeeded int
}

func (p *present) computeWrappings() {
	lw, wh, hl := p.length*p.width, p.width*p.height, p.height*p.length
	var sortedSurfaces []int
	sortedSurfaces = append(sortedSurfaces, lw, wh, hl)
	sort.Ints(sortedSurfaces)
	// you need 2 wrappings for each side, plus one wrapping more for the smallest area
	p.wrappingSurfaceNeeded = (sortedSurfaces[0] * 3) + (sortedSurfaces[1] * 2) + (sortedSurfaces[2] * 2)

	var sortedDims []int
	sortedDims = append(sortedDims, (p.length+p.width)*2, (p.width+p.height)*2, (p.height+p.length)*2)
	sort.Ints(sortedDims)
	// shortest distance around its sides, or the smallest perimeter of any one face
	// plus perfect bow is equal to the cubic feet of volume of the present
	p.ribbonNeeded = sortedDims[0] + (p.length * p.width * p.height)
}

func loadPresentsData(fileName string) *[]*present {
	readFile := utils.LoadInput(fileName)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	presents := make([]*present, 0, 10)

	for fileScanner.Scan() {
		currentDimensions := fileScanner.Text()
		dim := strings.Split(currentDimensions, "x")
		l, _ := strconv.Atoi(dim[0])
		w, _ := strconv.Atoi(dim[1])
		h, _ := strconv.Atoi(dim[2])
		present := &present{
			length:                l,
			width:                 w,
			height:                h,
			wrappingSurfaceNeeded: 0,
		}
		present.computeWrappings()
		presents = append(presents, present)
	}
	readFile.Close()

	return &presents
}
