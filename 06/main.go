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

	numbers, operations := getData()

	total := 0

	for i, operation := range operations {
		if operation == byte('+') {
			sum := 0
			for _, number := range numbers[i] {
				sum += number
			}
			total += sum
		} else {
			product := 1
			for _, number := range numbers[i] {
				product *= number
			}
			total += product
		}
	}

	fmt.Println("Total:", total)
	fmt.Println("Duration:", time.Since(start_time))
}

func getData() ([][]int, []byte) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var numbers [][]int
	var operations []byte

	var lines [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		lines = append(lines, parts)
	}
	for i := 0; i < len(lines[0]); i++ {
		numbers = append(numbers, []int{})
	}
	for i, line := range lines {
		if i == len(lines)-1 {
			for _, char := range line {
				operations = append(operations, byte(char[0]))
			}
		} else {
			for j, number_string := range line {
				number, _ := strconv.Atoi(number_string)
				numbers[j] = append(numbers[j], number)
			}
		}
	}

	return numbers, operations
}
