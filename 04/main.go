package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	start_time := time.Now()

	grid := getGrid()

	moveable_rolls := 0

	for y, row := range grid {
		for x, value := range row {
			if value == byte('.') {
				continue
			}

			adjacent_rolls := 0

			// Check above
			if y > 0 {
				if x > 0 && grid[y-1][x-1] == byte('@') {
					adjacent_rolls++
				}
				if grid[y-1][x] == byte('@') {
					adjacent_rolls++
				}
				if x < len(row)-1 && grid[y-1][x+1] == byte('@') {
					adjacent_rolls++
				}
			}

			// Check same line
			if x > 0 && grid[y][x-1] == byte('@') {
				adjacent_rolls++
			}
			if x < len(row)-1 && grid[y][x+1] == byte('@') {
				adjacent_rolls++
			}

			// Check below
			if y < len(grid)-1 {
				if x > 0 && grid[y+1][x-1] == byte('@') {
					adjacent_rolls++
				}
				if grid[y+1][x] == byte('@') {
					adjacent_rolls++
				}
				if x < len(row)-1 && grid[y+1][x+1] == byte('@') {
					adjacent_rolls++
				}
			}

			if adjacent_rolls < 4 {
				moveable_rolls++
			}
		}
	}

	fmt.Println("Moveable rolls:", moveable_rolls)
	fmt.Println("Duration:", time.Since(start_time))
}

func getGrid() [][]byte {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var numbers [][]byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers = append(numbers, []byte(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return numbers
}
