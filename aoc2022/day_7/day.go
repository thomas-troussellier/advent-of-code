package day_7

import (
	"advent-of-code/aoc"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type day struct {
	inputFile string
}

var _ aoc.EventRunner = (*day)(nil)

func Create() aoc.EventRunner {
	return create("aoc2022/day_7/input.txt")
}

func create(inputFile string) *day {
	return &day{
		inputFile: inputFile,
	}
}

func (d *day) Execute() {
	log.Println("Day 7")
	log.Println("Q1:", d.Question1())
	log.Println("Q2:", d.Question2())
}

func (d *day) Question1() string {
	tree := loadDirStruct(d.inputFile, 4)

	return strconv.Itoa(sumSizeAtMost(tree, 100000))
}

func (d *day) Question2() string {
	tree := loadDirStruct(d.inputFile, 4)

	return strconv.Itoa(chooseDirectoryToDelete(tree))
}

func chooseDirectoryToDelete(tree map[string]*content) int {
	// must have 30000000 free space on our 70000000 filesystem
	currentlyUsed := tree["/"].size
	free := 70000000 - currentlyUsed

	needed := 30000000 - free
	currentDeletedSize := 0

	return toDelete(tree, needed, currentDeletedSize)
}

func toDelete(tree map[string]*content, needed, deletedSize int) int {
	size := deletedSize
	for _, cont := range tree {
		if cont.content == nil {
			continue
		}

		if cont.size >= needed {
			if size == 0 || cont.size <= size {
				size = cont.size
			}
			size = toDelete(cont.content, needed, size)
		}

	}

	return size
}

func sumSizeAtMost(tree map[string]*content, maxSize int) int {
	var sum int
	for _, cont := range tree {
		if cont.content != nil {
			if maxSize >= cont.size {
				sum += cont.size
			}
			sum += sumSizeAtMost(cont.content, maxSize)
		}
	}
	return sum
}

type content struct {
	parent  *content
	content map[string]*content
	name    string
	size    int
}

func (c content) String() string {
	return fmt.Sprintf("%s, s:%d, \n%v", c.name, c.size, c.content)
}

func loadDirStruct(fileName string, diffChars int) map[string]*content {
	readFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	tree := make(map[string]*content)
	var currentDir *content

	for fileScanner.Scan() {
		currentLine := fileScanner.Text()

		words := strings.Split(currentLine, " ")

		switch words[0] {
		case "$":
			// matches a command
			switch words[1] {
			case "cd":
				switch words[2] {
				case "/":
					if _, ok := tree["/"]; !ok {
						tree["/"] = &content{
							name:    "/",
							size:    0,
							parent:  nil,
							content: make(map[string]*content),
						}
					}
					currentDir = tree["/"]
				case "..":
					currentDir = currentDir.parent
				default:
					currentDir = currentDir.content[words[2]]
				}
			case "ls":
			}
		case "dir":
			// matches a dir name
			currentDir.content[words[1]] = &content{
				name:    words[1],
				size:    0,
				parent:  currentDir,
				content: make(map[string]*content),
			}
		default:
			// matches a size
			size, _ := strconv.Atoi(words[0])
			currentDir.content[words[1]] = &content{
				name:    words[1],
				size:    size,
				parent:  currentDir,
				content: nil,
			}
			temp := currentDir
			for temp != nil {
				temp.size += size
				temp = temp.parent
			}
		}
	}

	readFile.Close()

	return tree
}
