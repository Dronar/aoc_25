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

	numbers := getNumbers()

	var power []int

	for _, number := range numbers {
		current_power := [2]byte{byte('0'), byte('0')}

		i := 0
		for i < len(number) {
			digit := number[i]

			if digit > current_power[0] && i+1 < len(number) {
				current_power[0] = digit
				current_power[1] = number[i+1]
				i++
				continue
			}
			if digit > current_power[1] {
				current_power[1] = digit
			}
			i++
		}
		temp, _ := strconv.Atoi(string(current_power[:]))

		power = append(power, temp)
	}

	sum := 0
	for _, num := range power {
		sum += num
	}

	fmt.Println("Sum:", sum)
	fmt.Println("Duration:", time.Since(start_time))
}

func getNumbers() [][]byte {
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
