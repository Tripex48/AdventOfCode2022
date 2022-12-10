package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func checkTop(grid [][]int, i int, j int) int {
	num := grid[i][j]
	score := 0
	for x := i - 1; x >= 0; x-- {
		if num <= grid[x][j] {
			score += 1
			break
		} else if num >= grid[x][j] {
			score += 1
		}
	}
	return score
}

func checkBottom(grid [][]int, i int, j int) int {
	num := grid[i][j]
	score := 0
	for x := i + 1; x < len(grid); x++ {
		if num <= grid[x][j] {
			score += 1
			break
		} else if num >= grid[x][j] {
			score += 1
		}
	}
	return score
}

func checkLeft(grid [][]int, i int, j int) int {
	num := grid[i][j]
	score := 0
	for x := j - 1; x >= 0; x-- {
		if num <= grid[i][x] {
			score += 1
			break
		} else if num >= grid[i][x] {
			score += 1
		}
	}
	return score
}

func checkRight(grid [][]int, i int, j int) int {
	num := grid[i][j]
	score := 0
	for x := j + 1; x < len(grid); x++ {
		if num <= grid[i][x] {
			score += 1
			break
		} else if num >= grid[i][x] {
			score += 1
		}
	}
	return score
}

func main() {
	fmt.Println("--- Running Day 8 ---")

	start := time.Now()

	file, err := os.ReadFile("./08.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Unable to read file")
		os.Exit(-1)
	}
	fileContent := string(file)

	lines := strings.Split(strings.TrimSpace(fileContent), "\n")
	grid := [][]int{}

	for _, line := range lines {
		nums := []int{}
		for _, x := range line {
			num, _ := strconv.Atoi(string(x))
			nums = append(nums, num)
		}
		grid = append(grid, nums)
	}

	maxScore := 0
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			score := checkTop(grid, i, j) * checkLeft(grid, i, j) * checkRight(grid, i, j) * checkBottom(grid, i, j)
			if score > maxScore {
				maxScore = score
			}
		}
	}

	fmt.Printf("Output Part 2: %d\n", maxScore)

	elapsed := time.Since(start)
	fmt.Printf("\nExecution took %s\n", elapsed)
}
