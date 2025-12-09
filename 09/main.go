package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	start_time := time.Now()

	tiles := getTiles()
	largest := 0

	for i, tile1 := range tiles {
		for j := i + 1; j < len(tiles); j++ {
			tile2 := tiles[j]

			length := tile2[0] - tile1[0]
			if length < 0 {
				length *= -1
			}
			height := tile2[1] - tile1[1]
			if height < 0 {
				height *= -1
			}

			area := (length + 1) * (height + 1)
			if area < 0 {
				area *= -1
			}

			if area > largest {
				largest = area
			}
		}
	}

	fmt.Println("Area:", largest)
	fmt.Println("Duration:", time.Since(start_time))
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

	sort.Slice(tiles, func(a, b int) bool {
		return tiles[a][0] < tiles[b][0]
	})

	return tiles
}
