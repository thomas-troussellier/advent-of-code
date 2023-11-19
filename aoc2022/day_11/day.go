package day_11

import (
	"advent-of-code/aoc"
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type day struct {
	inputFile string
}

var _ aoc.EventRunner = (*day)(nil)

func Create() aoc.EventRunner {
	return create("aoc2022/day_11/input.txt")
}

func create(inputFile string) *day {
	return &day{
		inputFile: inputFile,
	}
}

func (d *day) Execute() {
	log.Println("Day 11")
	log.Println("Q1:", d.Question1())
	log.Println("Q2:", d.Question2())
}

func (d *day) Question1() string {
	monkeys := loadMonkeys(d.inputFile)
	usage := playMonkeyBusiness(monkeys, 20, true)

	return strconv.FormatInt(monkeyBusiness(usage), 10)
}

func (d *day) Question2() string {
	monkeys := loadMonkeys(d.inputFile)
	usage2 := playMonkeyBusiness(monkeys, 10000, false)

	return strconv.FormatInt(monkeyBusiness(usage2), 10)
}

func monkeyBusiness(usage map[int]int64) int64 {
	uses := make([]int64, 0)

	for _, v := range usage {
		uses = append(uses, v)
	}
	sort.SliceStable(uses, func(i, j int) bool {
		return uses[i] < uses[j]
	})

	for left, right := 0, len(uses)-1; left < right; left, right = left+1, right-1 {
		uses[left], uses[right] = uses[right], uses[left]
	}
	fmt.Println(uses)
	return uses[0] * uses[1]
}

func playMonkeyBusiness(monkeys map[int]*monkey, rounds int, alleviateWorry bool) map[int]int64 {
	u := make(map[int]int64)

	for i := 0; i < rounds; i++ {
		roundMonkeyBusiness(monkeys, u, alleviateWorry, i)
	}

	return u
}

func roundMonkeyBusiness(monkeys map[int]*monkey, monkeyUse map[int]int64, alleviateWorry bool, round int) {
	for m := 0; m < len(monkeys); m++ {
		monkey := monkeys[m]

		items := monkey.startItems

		for _, worry := range items {
			wLvl := monkey.operation(worry)
			if alleviateWorry {
				wLvl = int64(math.Floor(float64(wLvl) / float64(3)))
			} else {
				// need to find appropriate transformation
			}

			test := monkey.test(wLvl)

			if test {
				monkeys[monkey.testTrue].startItems = append(monkeys[monkey.testTrue].startItems, wLvl)
			} else {
				monkeys[monkey.testFalse].startItems = append(monkeys[monkey.testFalse].startItems, wLvl)
			}

			monkey.startItems = monkey.startItems[1:]
			monkeyUse[monkey.name] = monkeyUse[monkey.name] + 1
		}
	}
}

type monkey struct {
	operation  func(int64) int64
	test       func(int64) bool
	startItems []int64
	name       int
	testTrue   int
	testFalse  int
}

func (m monkey) String() string {
	t := m.operation(10)
	return fmt.Sprintf("monkey %d, start: %v, mTrue: %d, mFalse: %d, op: %d", m.name, m.startItems, m.testTrue, m.testFalse, t)
}

func loadMonkeys(fileName string) map[int]*monkey {
	readFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	monkeys := make(map[int]*monkey, 0)

	var monkey monkey

	for fileScanner.Scan() {
		currentLine := fileScanner.Text()
		currentLine = strings.TrimSpace(currentLine)

		if strings.HasPrefix(currentLine, "Monkey") {
			var name int
			fmt.Sscanf(currentLine, "Monkey %d:", &name)
			monkey.name = name
		} else if strings.HasPrefix(currentLine, "Starting") {

			objects := strings.SplitAfter(currentLine, ":")
			objects = strings.Split(objects[1], ",")
			monkey.startItems = make([]int64, 0)
			for _, s := range objects {
				st := strings.TrimSpace(s)

				item, _ := strconv.Atoi(st)
				monkey.startItems = append(monkey.startItems, int64(item))
			}
		} else if strings.HasPrefix(currentLine, "Operation") {
			var op string
			var amountS string
			var amount int
			fmt.Sscanf(currentLine, "Operation: new = old %s %s", &op, &amountS)

			if amountS == "old" {
				op = amountS + op
			} else {
				amount, _ = strconv.Atoi(amountS)
			}
			monkey.operation = operation(op, amount)
		} else if strings.HasPrefix(currentLine, "Test") {
			var divisible int
			fmt.Sscanf(currentLine, "Test: divisible by %d", &divisible)
			monkey.test = testFnc(divisible)
		} else if strings.HasPrefix(currentLine, "If") {
			s := strings.Split(currentLine, " ")
			monk, _ := strconv.Atoi(s[len(s)-1])
			if s[1][:len(s[1])-1] == "true" {
				monkey.testTrue = monk
			} else {
				monkey.testFalse = monk
			}
		} else {
			var m = monkey
			monkeys[monkey.name] = &m
		}
	}
	monkeys[monkey.name] = &monkey

	readFile.Close()

	return monkeys
}

func operation(operation string, supplement int) func(int64) int64 {
	if operation == "*" {
		return func(worry int64) int64 {
			return worry * int64(supplement)
		}
	}
	if operation == "+" {
		return func(worry int64) int64 {
			return worry + int64(supplement)
		}
	}
	if operation == "old+" {
		return func(worry int64) int64 {
			return worry + worry
		}
	}
	if operation == "old*" {
		return func(worry int64) int64 {
			return worry * worry
		}
	}

	return func(worry int64) int64 {
		return worry
	}
}

func testFnc(div int) func(int64) bool {
	return func(worry int64) bool {
		return math.Mod(float64(worry), float64(div)) == 0
	}
}
