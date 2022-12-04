package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func contains(self_start int, self_end int, other_start int, other_end int) bool {
	return self_start <= other_start && self_end >= other_end
}

func overlap(self_start int, self_end int, other_start int, other_end int) bool {
	if self_end < other_start || other_end < self_start {
		return false
	}
	return true
}

func main() {
	fmt.Println("--- Running Day 4 ---")

	start := time.Now()

	file, err := os.ReadFile("./04.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Unable to read file")
		os.Exit(-1)
	}
	fileContent := string(file)

	input := strings.Split(strings.TrimSpace(fileContent), "\n")

	count_p1 := 0
	count_p2 := 0
	for _, line := range input {
		parts := strings.Split(line, ",")
		p1_nums := ParseNumbers(strings.Split(parts[0], "-"))
		p2_nums := ParseNumbers(strings.Split(parts[1], "-"))

		if contains(p1_nums[0], p1_nums[1], p2_nums[0], p2_nums[1]) || contains(p2_nums[0], p2_nums[1], p1_nums[0], p1_nums[1]) {
			count_p1 += 1
		}

		if overlap(p1_nums[0], p1_nums[1], p2_nums[0], p2_nums[1]) {
			count_p2 += 1
		}
	}

	fmt.Println(fmt.Sprintf("Output Part 1: %d", count_p1))
	fmt.Println(fmt.Sprintf("Output Part 2: %d", count_p2))

	elapsed := time.Since(start)
	fmt.Printf("\nExecution took %s\n", elapsed)
}

// ParseNumbers returns a list of int from a list of string
func ParseNumbers(lines []string) []int {
	values := []int{}

	for _, line := range lines {
		value, _ := strconv.Atoi(line)
		values = append(values, value)
	}

	return values
}
