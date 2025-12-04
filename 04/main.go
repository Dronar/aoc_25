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

	moveable_rolls := step2(grid)

	fmt.Println("Moveable rolls:", moveable_rolls)
	fmt.Println("Duration:", time.Since(start_time))
}

func step1(grid [][]byte) int {
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

	return moveable_rolls
}

func step2(grid [][]byte) int {
	moveable_rolls := 0

	moved_this_round := true

	for moved_this_round {
		var rolls_to_remove [][2]int
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
					rolls_to_remove = append(rolls_to_remove, [2]int{y, x})
				}
			}
		}
		for _, coordinates := range rolls_to_remove {
			grid[coordinates[0]][coordinates[1]] = byte('.')
		}
		moveable_rolls += len(rolls_to_remove)
		moved_this_round = len(rolls_to_remove) > 0
	}

	return moveable_rolls
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
