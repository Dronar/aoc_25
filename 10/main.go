package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mowshon/iterium"
)

func main() {
	start_time := time.Now()

	machines, buttons := getSchema()

	button_clicks := 0
	for i, machine := range machines {
		button_list := buttons[i]

		for j := range 10 {
			clicks := 0
			combinations := iterium.CombinationsWithReplacement(button_list, j)
			for {
				val, err := combinations.Next()
				if err != nil {
					break
				}
				sum := 0
				for _, v := range val {
					sum ^= v
				}
				if sum == machine {
					clicks = j
					break
				}
			}
			if clicks > 0 {
				button_clicks += clicks
				break
			}
		}
	}

	fmt.Println("Clicks:", button_clicks)
	fmt.Println("Duration:", time.Since(start_time))
}

func getSchema() ([]int, [][]int) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var machines []int
	var buttons [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		split_line := strings.Split(line, "] ")

		// Get machine
		machine := split_line[0][1:]
		machine = strings.ReplaceAll(machine, "#", "1")
		machine = strings.ReplaceAll(machine, ".", "0")
		binary, _ := strconv.ParseInt(machine, 2, 64)
		machines = append(machines, int(binary))

		parts := strings.Split(split_line[1], ") {")

		// Get buttons
		button_split := parts[0][1:]
		button_parts := strings.Split(button_split, ") (")

		var button_numbers []int

		for _, button := range button_parts {
			var button_string []rune
			for i := 0; i < len(machine); i++ {
				button_string = append(button_string, '0')
			}
			nums := strings.Split(button, ",")

			for _, num := range nums {
				button_number, _ := strconv.Atoi(num)
				button_string[button_number] = '1'
			}
			value, _ := strconv.ParseInt(string(button_string), 2, 64)
			button_numbers = append(button_numbers, int(value))
		}
		buttons = append(buttons, button_numbers)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return machines, buttons
}
