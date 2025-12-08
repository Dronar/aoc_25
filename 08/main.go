package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"
)

type point struct {
	x int
	y int
	z int
}

type distance struct {
	length int
	p1     point
	p2     point
}

func main() {
	start_time := time.Now()

	points := getPoints()

	var distances []distance
	for i := range len(points) {
		for j := i + 1; j < len(points); j++ {
			p1 := points[i]
			p2 := points[j]
			dist := distance{
				length: getDistance(points[i], points[j]),
				p1:     p1,
				p2:     p2,
			}

			distances = append(distances, dist)
		}
	}

	sort.Slice(distances, func(a, b int) bool {
		return distances[a].length < distances[b].length
	})

	step2(distances)

	fmt.Println("Duration:", time.Since(start_time))
}

func step1(distances []distance) {
	var connections [][]point

	for i := range 1000 {
		found1 := -1
		found2 := -1

		for j, connection := range connections {
			if slices.Contains(connection, distances[i].p1) && found1 < 0 {
				found1 = j
			}
			if slices.Contains(connection, distances[i].p2) && found2 < 0 {
				found2 = j
			}
		}

		if found1 >= 0 && found1 == found2 {
			continue
		}
		if found1 >= 0 && found2 >= 0 {
			connections[found1] = append(connections[found1], connections[found2]...)
			connections[found2] = []point{}
			continue
		}
		if found1 >= 0 {
			connections[found1] = append(connections[found1], distances[i].p2)
			continue
		}
		if found2 >= 0 {
			connections[found2] = append(connections[found2], distances[i].p1)
			continue
		}

		connections = append(connections, []point{})
		connections[len(connections)-1] = append(connections[len(connections)-1], distances[i].p1)
		connections[len(connections)-1] = append(connections[len(connections)-1], distances[i].p2)
	}

	largest := 0
	second := 0
	third := 0
	for _, con := range connections {
		if len(con) > largest {
			third = second
			second = largest
			largest = len(con)
			continue
		}
		if len(con) > second {
			third = second
			second = len(con)
			continue
		}
		if len(con) > third {
			third = len(con)
		}
	}

	fmt.Println("Circuits:", largest*second*third)
}

func step2(distances []distance) {
	var connections [][]point

	for i, _ := range distances {
		found1 := -1
		found2 := -1

		for j, connection := range connections {
			if slices.Contains(connection, distances[i].p1) && found1 < 0 {
				found1 = j
			}
			if slices.Contains(connection, distances[i].p2) && found2 < 0 {
				found2 = j
			}
		}

		if found1 >= 0 && found1 == found2 {
		} else if found1 >= 0 && found2 >= 0 {
			connections[found1] = append(connections[found1], connections[found2]...)
			connections[found2] = connections[len(connections)-1]
			connections = connections[:len(connections)-1]

			if i > 100 && len(connections) == 1 {
				fmt.Printf("Circuit complete (iteration: %d)! ", i)
				fmt.Printf("[%d, %d, %d] - ", distances[i].p1.x, distances[i].p1.y, distances[i].p1.z)
				fmt.Printf("[%d, %d, %d]\n", distances[i].p2.x, distances[i].p2.y, distances[i].p2.z)
				fmt.Println("Sum: ", distances[i].p1.x*distances[i].p2.x)

				break
			}

		} else if found1 >= 0 {
			connections[found1] = append(connections[found1], distances[i].p2)
		} else if found2 >= 0 {
			connections[found2] = append(connections[found2], distances[i].p1)
		} else {
			connections = append(connections, []point{})
			connections[len(connections)-1] = append(connections[len(connections)-1], distances[i].p1)
			connections[len(connections)-1] = append(connections[len(connections)-1], distances[i].p2)
		}
	}
}

func getDistance(p1, p2 point) int {
	d := math.Pow(math.Abs(float64(p1.x-p2.x)), 2) + math.Pow(math.Abs(float64(p1.y-p2.y)), 2) + math.Pow(math.Abs(float64(p1.z-p2.z)), 2)
	return int(math.Trunc(math.Sqrt(float64(d))))
}

func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}
	if m == 1 {
		return n
	}
	result := n
	for i := 2; i < m; i++ {
		result *= n
	}
	return n
}

func getPoints() []point {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var points []point

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number_string := scanner.Text()
		numbers := strings.Split(number_string, ",")

		var xyz []int
		for _, number := range numbers {
			num, _ := strconv.Atoi(number)
			xyz = append(xyz, num)
		}

		p := point{x: xyz[0], y: xyz[1], z: xyz[2]}

		points = append(points, p)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return points
}
