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

	// numbers, operations := step1()
	numbers, operations := step2()

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

func step1() ([][]int, []byte) {
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

func step2() ([][]int, []byte) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var operations []byte
	var numbers [][]int
	numbers = append(numbers, []int{})

	var grid [][]byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}

	column := len(grid[0]) - 1

	for column >= 0 {
		row := 0
		var digits []byte

		for row < len(grid) {
			value := grid[row][column]
			if value >= byte('1') && value <= byte('9') {
				digits = append(digits, value)
			}
			if value == byte('+') || value == byte('*') {
				operations = append(operations, value)
			}
			row++
		}

		if len(digits) > 0 {
			number, _ := strconv.Atoi(string(digits[:]))
			numbers[len(numbers)-1] = append(numbers[len(numbers)-1], number)
		} else {
			numbers = append(numbers, []int{})
		}
		column--
	}

	return numbers, operations
}
