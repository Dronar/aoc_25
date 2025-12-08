package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	start_time := time.Now()
	hits := step2()
	fmt.Println("Hits:", hits)
	fmt.Println("Duration:", time.Since(start_time))
}

func step1() int {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	counter := 0

	scanner := bufio.NewScanner(file)
	var grid [][]byte
	i := 0

	for scanner.Scan() {
		row := []byte(scanner.Text())

		if i == 1 {
			for j, _ := range row {
				if grid[0][j] == byte('S') {
					row[j] = byte('|')
				}
			}
		}
		if i > 1 {
			for j, value := range row {
				if grid[i-1][j] == byte('|') {
					if value == byte('^') {
						counter++
						row[j-1] = byte('|')
						row[j+1] = byte('|')
					} else {
						row[j] = byte('|')
					}
				}
			}
		}

		grid = append(grid, row)
		i++
	}

	return counter
}

func step2() int {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]int
	i := 0

	for scanner.Scan() {
		line := []byte(scanner.Text())
		var row []int

		j := 0
		for j < len(line) {
			value := line[j]
			if value == byte('.') {
				if i > 0 && grid[i-1][j] > 0 {
					row = append(row, grid[i-1][j])
				} else {
					row = append(row, 0)
				}
			}
			if value == byte('S') {
				row = append(row, 1)
			}
			if value == byte('^') {
				beams := grid[i-1][j]
				if beams > 0 {
					row[j-1] += beams
					row = append(row, 0)
					row = append(row, beams+grid[i-1][j+1])
					j++
				} else {
					row = append(row, 0)
				}
			}

			j++
		}

		grid = append(grid, row)
		i++
	}

	counter := 0
	for _, beams := range grid[len(grid)-1] {
		counter += beams
	}
	return counter
}
