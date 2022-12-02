package aoc2015

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func day2() {
	fmt.Println("Day 2")

	// load exercise data
	presents := LoadData("aoc2015/day_2.log")

	// how many square feet of wrapping paper ?
	totalWrapping := 0
	totalRibbon := 0
	for _, p := range *presents {
		totalWrapping += p.wrappingSurfaceNeeded
		totalRibbon += p.ribbonNeeded
	}
	fmt.Println("Q1. How many total square feet of wrapping paper should they order?")
	fmt.Printf("A1. total : %d\n", totalWrapping)

	fmt.Println("Q2. How many total feet of ribbon should they order?")
	fmt.Printf("A2. total : %d\n", totalRibbon)
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

func LoadData(fileName string) *[]*present {
	readFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
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
