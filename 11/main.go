package main

import (
	"bufio"
	"fmt"
	"os"
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

func (g *Graph) DFS(node string, target string, visited map[string]bool, sum *int) {
	if node == target {
		*sum++
		return
	}
	visited[node] = true

	for _, w := range g.vertices[node] {
		if !visited[w] {
			g.DFS(w, target, visited, sum)
		}
	}

	visited[node] = false
}

func main() {
	start_time := time.Now()

	connections := getConnections()

	valid_paths := 0

	visited := make(map[string]bool)
	connections.DFS("you", "out", visited, &valid_paths)

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
