package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Coord struct {
	x int
	y int
}

func main() {
	fmt.Println("--- Running Day 9 ---")

	start := time.Now()

	file, err := os.ReadFile("./09.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Unable to read file")
		os.Exit(-1)
	}
	fileContent := string(file)

	lines := strings.Split(strings.TrimSpace(fileContent), "\n")

	visited := moveSnake(lines, 2)
	fmt.Printf("Output Part 1: %d\n", len(visited))

	visited = moveSnake(lines, 10)
	fmt.Printf("Output Part 2: %d\n", len(visited))

	elapsed := time.Since(start)
	fmt.Printf("\nExecution took %s\n", elapsed)
}

func moveSnake(lines []string, length int) map[Coord]struct{} {
	visited := make(map[Coord]struct{})

	snake := [][]int{}
	for i := 0; i < length; i++ {
		temp := []int{0, 0}
		snake = append(snake, temp)
	}

	for _, line := range lines {
		fields := strings.Split(line, " ")
		x := fields[0]
		y, _ := strconv.Atoi(fields[1])

		for i := 0; i < y; i++ {
			dx := 0
			if x == "R" {
				dx = 1
			} else if x == "L" {
				dx = -1
			}

			dy := 0
			if x == "U" {
				dy = 1
			} else if x == "D" {
				dy = -1
			}

			snake[0][0] += dx
			snake[0][1] += dy

			for i := 0; i < length-1; i++ {
				H := snake[i]
				T := snake[i+1]

				_x := H[0] - T[0]
				_y := H[1] - T[1]

				if math.Abs(float64(_x)) > 1 || math.Abs(float64(_y)) > 1 {
					if _x == 0 {
						if _y > 0 {
							T[1] += 1
						} else {
							T[1] += -1
						}
					} else if _y == 0 {
						if _x > 0 {
							T[0] += 1
						} else {
							T[0] += -1
						}
					} else {
						if _x > 0 {
							T[0] += 1
						} else {
							T[0] += -1
						}

						if _y > 0 {
							T[1] += 1
						} else {
							T[1] += -1
						}
					}
				}
			}

			visited[Coord{x: snake[len(snake)-1][0], y: snake[len(snake)-1][1]}] = struct{}{}
		}
	}
	return visited
}
