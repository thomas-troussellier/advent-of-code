package day_12

import (
	"advent-of-code/aoc"
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

type day struct {
	inputFile string
}

var _ aoc.EventRunner = (*day)(nil)

func Create() aoc.EventRunner {
	return create("aoc2022/day_12/input.txt")
}

func create(inputFile string) *day {
	return &day{
		inputFile: inputFile,
	}
}

func (d *day) Execute() {
	log.Println("Day 12")
	log.Println("Q1:", d.Question1())
	log.Println("Q2:", d.Question2())
}

func (d *day) Question1() string {
	baseHeightMap := loadHeightMap(d.inputFile)

	graph, nodeS, nodeE, _ := buildDifferent(baseHeightMap)

	visited := shortestBis(graph, nodeS, nodeE, len(graph))

	return strconv.Itoa(len(visited) - 1)
}

func (d *day) Question2() string {

	baseHeightMap := loadHeightMap(d.inputFile)

	graph, _, nodeE, possStarts := buildDifferent(baseHeightMap)
	best := 0
	for _, start := range possStarts {
		res := shortestBis(graph, start, nodeE, len(graph))
		if len(res) == 0 {
			continue
		}
		if best == 0 || len(res) < best {
			best = len(res)
		}
	}

	return strconv.Itoa(best - 1)
}

func shortestBis(adj map[string][]string, start string, end string, nbPnt int) []string {
	predecessor := make(map[string]string, nbPnt)
	dist := make(map[string]int, nbPnt)

	if !breadthFirstSearch(adj, start, end, nbPnt, predecessor, dist) {
		return []string{}
	}

	var shortest []string
	crawl := end
	shortest = append(shortest, crawl)

	for predecessor[crawl] != "" {
		shortest = append(shortest, predecessor[crawl])
		crawl = predecessor[crawl]
	}

	return shortest
}

func breadthFirstSearch(adj map[string][]string, start, end string, nbPnt int, predecessor map[string]string, dist map[string]int) bool {
	var (
		queue   []string
		visited = make(map[string]bool)
	)

	for k := range adj {
		visited[k] = false
		dist[k] = math.MaxInt
		predecessor[k] = ""
	}

	visited[start] = true
	dist[start] = 0
	queue = append(queue, start)

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		for i := 0; i < len(adj[u]); i++ {
			if !visited[adj[u][i]] {
				visited[adj[u][i]] = true
				dist[adj[u][i]] = dist[u] + 1
				predecessor[adj[u][i]] = u
				queue = append(queue, adj[u][i])

				if adj[u][i] == end {
					return true
				}
			}
		}
	}

	return false
}

func buildDifferent(baseHeightMap []string) (graph map[string][]string, start string, end string, possibleStarts []string) {
	graph = make(map[string][]string)
	possibleStarts = make([]string, 0)

	for i, heightRow := range baseHeightMap {
		for j, heightColumn := range heightRow {
			height := string(heightColumn)
			cx := strconv.Itoa(i)
			cy := strconv.Itoa(j)
			name := cx + "," + cy
			graph[name] = make([]string, 0)

			if i != 0 {
				x := strconv.Itoa(i - 1)
				upName := x + "," + cy
				h := string(baseHeightMap[i-1][j])
				switch h {
				case "S":
					h = "a"
				case "E":
					h = "z"
				}
				if reachable(h, height) {
					graph[upName] = append(graph[upName], name)
				}
				if reachable(height, h) {
					graph[name] = append(graph[name], upName)
				}
			}

			if j != 0 {
				y := strconv.Itoa(j - 1)
				leftName := cx + "," + y
				h := string(baseHeightMap[i][j-1])
				switch h {
				case "S":
					h = "a"
				case "E":
					h = "z"
				}
				if reachable(h, height) {
					graph[leftName] = append(graph[leftName], name)
				}
				if reachable(height, h) {
					graph[name] = append(graph[name], leftName)
				}
			}

			if height == "S" {
				start = name
			}

			if height == "E" {
				end = name
			}

			if height == "S" || height == "a" {
				possibleStarts = append(possibleStarts, name)
			}
		}
	}

	return graph, start, end, possibleStarts
}

func reachable(from, to string) bool {
	if to == "S" {
		return false
	}

	if to == "E" && (from == "z" || from == "y") {
		return true
	} else if to == "E" {
		return false
	}

	if from == "S" && (to == "a" || to == "b") {
		return true
	}

	ifrom := from[0]
	ito := to[0]

	return ito == ifrom+1 || ifrom >= ito
}

func loadHeightMap(fileName string) []string {
	readFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	lines := make([]string, 0)

	for fileScanner.Scan() {
		currentLine := fileScanner.Text()

		lines = append(lines, currentLine)
	}

	readFile.Close()

	return lines
}
