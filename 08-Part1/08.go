package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func checkTop(grid [][]int, i int, j int) bool {
	num := grid[i][j]
	for x := 0; x < i; x++ {
		if num <= grid[x][j] {
			return false
		}
	}
	return true
}

func checkBottom(grid [][]int, i int, j int) bool {
	num := grid[i][j]
	for x := i + 1; x < len(grid); x++ {
		if num <= grid[x][j] {
			return false
		}
	}
	return true
}

func checkLeft(grid [][]int, i int, j int) bool {
	num := grid[i][j]
	for x := 0; x < j; x++ {
		if num <= grid[i][x] {
			return false
		}
	}
	return true
}

func checkRight(grid [][]int, i int, j int) bool {
	num := grid[i][j]
	for x := j + 1; x < len(grid); x++ {
		if num <= grid[i][x] {
			return false
		}
	}
	return true
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
	visible := len(lines[0])*2 + (len(lines) * 2) - 4
	fmt.Println(visible)

	for _, line := range lines {
		nums := []int{}
		for _, x := range line {
			num, _ := strconv.Atoi(string(x))
			nums = append(nums, num)
		}
		grid = append(grid, nums)
	}

	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			if checkTop(grid, i, j) || checkLeft(grid, i, j) || checkRight(grid, i, j) || checkBottom(grid, i, j) {
				visible += 1
			}
		}
	}

	fmt.Printf("Output Part 1: %d\n", visible)

	elapsed := time.Since(start)
	fmt.Printf("\nExecution took %s\n", elapsed)
}
