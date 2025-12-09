package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start_time := time.Now()

	tiles := getTiles()
	//area := step1(tiles)
	area := step2(tiles)

	fmt.Println("Area:", area)
	fmt.Println("Duration:", time.Since(start_time))
}

func step1(tiles [][2]int) int {
	largest := 0

	for i, tile1 := range tiles {
		for j := i + 1; j < len(tiles); j++ {
			tile2 := tiles[j]

			area := calculate_area(tile1, tile2)

			if area < 0 {
				area *= -1
			}

			if area > largest {
				largest = area
			}
		}
	}

	return largest
}

func step2(tiles [][2]int) int {
	var grid [100000][100000]bool

	// Create grid
	for i, tile := range tiles {
		var next [2]int
		if i < len(tiles)-1 {
			next = tiles[i+1]
		} else {
			next = tiles[0]
		}

		if tile[0] == next[0] {
			start := min(tile[1], next[1])
			end := max(tile[1], next[1])

			for j := start; j <= end; j++ {
				grid[j][tile[0]] = true
			}
		} else {
			start := min(tile[0], next[0])
			end := max(tile[0], next[0])

			for j := start; j <= end; j++ {
				grid[tile[1]][j] = true
			}
		}
	}

	// Fill grid
	for i, _ := range grid {
		start := 0
		for j, val := range grid[i] {
			if start == 0 && val {
				start = j
				continue
			}
			if start > 0 && val {
				for k := start; k < j; k++ {
					grid[i][k] = true
				}
				start = j
			}
		}
	}

	// Find rectangles

	largest := 0

	for i, tile1 := range tiles {
		for j := i + 1; j < len(tiles); j++ {
			tile2 := tiles[j]

			area := calculate_area(tile1, tile2)

			if area < 0 {
				area *= -1
			}

			if area < largest {
				continue
			}

			// Check borders
			in_range := true
			for k := min(tile1[0], tile2[0]); k <= max(tile1[0], tile2[0]); k++ {
				if !grid[tile1[1]][k] || !grid[tile2[1]][k] {
					in_range = false
					break
				}
			}
			if !in_range {
				continue
			}
			for k := min(tile1[1], tile2[1]); k <= max(tile1[1], tile2[1]); k++ {
				if !grid[k][tile1[0]] || !grid[k][tile2[0]] {
					in_range = false
					break
				}
			}
			if !in_range {
				continue
			}
			largest = area
		}
	}

	return largest
}

func calculate_area(tile1, tile2 [2]int) int {
	length := tile2[0] - tile1[0]
	if length < 0 {
		length *= -1
	}
	height := tile2[1] - tile1[1]
	if height < 0 {
		height *= -1
	}

	return (length + 1) * (height + 1)
}

func getTiles() [][2]int {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var tiles [][2]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		numbers := strings.Split(line, ",")

		x, _ := strconv.Atoi(numbers[0])
		y, _ := strconv.Atoi(numbers[1])

		tiles = append(tiles, [2]int{x, y})
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return tiles
}
