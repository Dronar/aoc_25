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

	fresh_ranges, _ := getData()

	//fresh_count := step1(fresh_ranges, ingredients)
	fresh_count := step2(fresh_ranges)

	fmt.Println("Fresh ingredients:", fresh_count)
	fmt.Println("Duration:", time.Since(start_time))
}

func step1(fresh_ranges [][2]int, ingredients []int) int {
	fresh_count := 0

	for _, item := range ingredients {
		for _, fresh_range := range fresh_ranges {
			if item > fresh_range[0] && item < fresh_range[1] {
				fresh_count++
				break
			}
		}
	}

	return fresh_count
}

func step2(fresh_ranges [][2]int) int {
	var union_ranges [][2]int
	for _, fresh_range := range fresh_ranges {
		should_insert := true
		var ranges_to_merge []int

		//fmt.Printf("Range: %d\n", fresh_range)
		for i, unique_range := range union_ranges {
			// If new is inside existing range
			if fresh_range[0] >= unique_range[0] && fresh_range[1] <= unique_range[1] {
				ranges_to_merge = append(ranges_to_merge, i)
				should_insert = false
				continue
			}
			// If existing range is inside new
			if fresh_range[0] <= unique_range[0] && fresh_range[1] >= unique_range[1] {
				ranges_to_merge = append(ranges_to_merge, i)
				should_insert = false
				continue
			}
			// If new is partly inside range (higher)
			if fresh_range[0] >= unique_range[0] && fresh_range[0] <= unique_range[1] {
				ranges_to_merge = append(ranges_to_merge, i)
				should_insert = false
				continue
			}
			// If new is partly inside range (lower)
			if fresh_range[1] >= unique_range[0] && fresh_range[1] <= unique_range[1] {
				ranges_to_merge = append(ranges_to_merge, i)
				should_insert = false
				continue
			}
		}

		if len(ranges_to_merge) > 0 {
			// Dummy copy to iterate over
			range_copy := append([][2]int(nil), union_ranges...)

			min := fresh_range[0]
			max := fresh_range[1]

			for _, r := range ranges_to_merge {
				if range_copy[r][0] < min {
					min = range_copy[r][0]
				}
				if range_copy[r][1] > max {
					max = range_copy[r][1]
				}

				// Delete range in original array
				union_ranges[r] = union_ranges[len(union_ranges)-1]
				union_ranges = union_ranges[:len(union_ranges)-1]
			}

			union_ranges = append(union_ranges, [2]int{min, max})
		}
		if should_insert {
			union_ranges = append(union_ranges, fresh_range)
		}
	}

	ingredient_count := 0

	for _, unique_range := range union_ranges {
		ingredient_count += unique_range[1] - unique_range[0] + 1
	}

	return ingredient_count
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
