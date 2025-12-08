package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	start_time := time.Now()
	hits := calculate()
	fmt.Println("Hits:", hits)
	fmt.Println("Duration:", time.Since(start_time))
}

func calculate() int {
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
