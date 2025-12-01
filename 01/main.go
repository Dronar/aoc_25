package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start_time := time.Now()
	//step_1()
	step_2()
	fmt.Println("Duration:", time.Since(start_time))
}

func step_1() {
	numbers := getNumbers()

	current := 50
	counter := 0

	for _, v := range numbers {
		current += v % 100
		if current > 99 {
			current = current % 100
		}
		if current < 0 {
			current = 100 + current
		}

		if current == 0 {
			counter++
		}
	}

	fmt.Println("Answer:", counter)
}

func step_2() {
	numbers := getNumbers()

	current := 50
	counter := 0

	for _, v := range numbers {
		current += v

		if current > 99 {
			for current > 99 {
				current -= 100
				counter++
			}
			if current == 0 {
				counter--
			}
		}

		if current < 0 {
			for current < 0 {
				current = 100 + current
				counter++
			}
			if (current-v)%100 == 0 {
				counter--
			}
		}

		if current == 0 {
			counter++
		}
	}

	fmt.Println("Answer:", counter)
}

func getNumbers() []int {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var numbers = make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number_string := scanner.Text()
		number, _ := strconv.Atoi(number_string[1:])

		if number_string[0] == 'R' {
			numbers = append(numbers, number)
		} else {
			numbers = append(numbers, 0-number)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return numbers
}
