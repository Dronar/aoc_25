package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start_time := time.Now()

	ranges := getRanges()

	var hits []int

	for _, r := range ranges {
		//fmt.Printf("%d - %d\n", r[0], r[1])
		i := r[0]

		for i <= r[1] {
			str := strconv.Itoa(i)
			len := len(str)

			if len%2 == 0 {
				part1 := str[:len/2]
				part2 := str[len/2:]

				if part1 == part2 {
					hits = append(hits, i)
				}

				i++
			} else {
				i = int(math.Pow10(len))
			}
		}

	}

	sum := 0
	for _, num := range hits {
		sum += num
	}

	//fmt.Printf("%d\n", hits)

	fmt.Println("Sum:", sum)
	fmt.Println("Duration:", time.Since(start_time))
}

func getRanges() [][2]int {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Err")
	}

	string_ranges := strings.Split(string(content), ",")

	var ranges [][2]int

	for _, r := range string_ranges {
		numbers := strings.Split(r, "-")

		start, _ := strconv.Atoi(numbers[0])
		end, _ := strconv.Atoi(numbers[1])

		var new_range [2]int

		new_range = [2]int{start, end}
		ranges = append(ranges, new_range)
	}

	return ranges
}
