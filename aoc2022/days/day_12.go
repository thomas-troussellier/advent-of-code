package days

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func Day12() {
	fmt.Println("Day 12")

	baseHeightMap := loadHeightMap("aoc2022/inputs/day_12.log")

	graph, nodeS, nodeE, possStarts := buildDifferent(baseHeightMap)

	visited := shortestBis(graph, nodeS, nodeE, len(graph))
	fmt.Println("Q1. What is the fewest steps required to move from your current position to the location that should get the best signal?")
	// don't count the starting S
	fmt.Printf("A1. fewest steps: %d\n", len(visited)-1)

	fmt.Println("Q2. What is the fewest steps required to move starting from any square with elevation a to the location that should get the best signal?")
	// don't count the starting S
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
	fmt.Printf("A2. fewest steps: %d\n", best-1)
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
