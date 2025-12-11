package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Graph struct {
	vertices map[string][]string
}

func (g *Graph) AddEdge(v, w string) {
	g.vertices[v] = append(g.vertices[v], w)
}
func NewGraph() *Graph {
	return &Graph{vertices: make(map[string][]string)}
}

func (g *Graph) DFS_part1(node string, target string, visited map[string]bool, sum *int) {
	if node == target {
		*sum++
		return
	}
	visited[node] = true

	for _, w := range g.vertices[node] {
		if !visited[w] {
			g.DFS_part1(w, target, visited, sum)
		}
	}

	visited[node] = false
}

func (g *Graph) DFS_part2(node string, target string, fft bool, dac bool, visited map[string]int) int {
	cache_key := node + strconv.FormatBool(fft) + strconv.FormatBool(dac)
	value, ok := visited[cache_key]
	if ok {
		return value
	}
	if node == target {
		if fft && dac {
			return 1
		} else {
			return 0
		}
	}

	if node == "fft" {
		fft = true
	}
	if node == "dac" {
		dac = true
	}

	paths := 0

	for _, neighbor := range g.vertices[node] {

		paths += g.DFS_part2(neighbor, target, fft, dac, visited)
	}

	visited[cache_key] = paths

	return paths
}

func main() {
	start_time := time.Now()

	connections := getConnections()

	valid_paths := 0

	//visited1 := make(map[string]bool)
	//connections.DFS_part1("svr", "out", visited1, &valid_paths)

	visited2 := make(map[string]int)
	valid_paths = connections.DFS_part2("svr", "out", false, false, visited2)

	fmt.Println("Paths:", valid_paths)
	fmt.Println("Duration:", time.Since(start_time))
}

func getConnections() Graph {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	connections := NewGraph()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ":")
		targets := strings.Fields(parts[1])

		for _, target := range targets {
			connections.AddEdge(parts[0], target)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return *connections
}
