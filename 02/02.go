package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("--- Running Day 2 ---")

	start := time.Now()

	file, err := os.ReadFile("./02.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Unable to read file")
		os.Exit(-1)
	}
	fileContent := string(file)

	input := strings.Split(fileContent, "\n")

	results := make(map[string]int)
	results["A X"] = 4
	results["A Y"] = 8
	results["A Z"] = 3
	results["B X"] = 1
	results["B Y"] = 5
	results["B Z"] = 9
	results["C X"] = 7
	results["C Y"] = 2
	results["C Z"] = 6

	total_score := play(input, results)
	fmt.Println(fmt.Sprintf("Output Part 1: %d", total_score))

	results_r2 := make(map[string]int)
	results_r2["A X"] = 3
	results_r2["A Y"] = 4
	results_r2["A Z"] = 8
	results_r2["B X"] = 1
	results_r2["B Y"] = 5
	results_r2["B Z"] = 9
	results_r2["C X"] = 2
	results_r2["C Y"] = 6
	results_r2["C Z"] = 7

	total_score_r2 := play(input, results_r2)
	fmt.Println(fmt.Sprintf("Output Part 2: %d", total_score_r2))

	elapsed := time.Since(start)
	fmt.Printf("\nExecution took %s\n", elapsed)
}

func play(input []string, results map[string]int) int {
	total_score := 0

	for _, line := range input {
		total_score += results[line]
	}

	return total_score
}
