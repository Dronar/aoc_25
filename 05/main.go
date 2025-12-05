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

	fresh_ranges, ingredients := getData()
	fresh_count := 0

	for _, item := range ingredients {
		for _, fresh_range := range fresh_ranges {
			if item > fresh_range[0] && item < fresh_range[1] {
				fresh_count++
				break
			}
		}
	}

	fmt.Println("Fresh ingredients:", fresh_count)
	fmt.Println("Duration:", time.Since(start_time))
}

func getData() ([][2]int, []int) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var fresh_ranges [][2]int
	var ingredients []int

	reading_ranges := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			reading_ranges = false
			continue
		}

		if reading_ranges {
			numbers := strings.Split(line, "-")
			start, _ := strconv.Atoi(numbers[0])
			end, _ := strconv.Atoi(numbers[1])

			fresh_ranges = append(fresh_ranges, [2]int{start, end})
		} else {
			number, _ := strconv.Atoi(line)
			ingredients = append(ingredients, number)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return fresh_ranges, ingredients
}
