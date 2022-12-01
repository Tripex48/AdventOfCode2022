package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("--- Running Day 1 ---")

	start := time.Now()

	file, err := os.ReadFile("./01.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Unable to read file")
		os.Exit(-1)
	}
	fileContent := string(file)

	input := strings.Split(fileContent, "\n\n")
	highest_elf := 0
	var elf_totals []int
	for _, elf_input := range input {
		nums := ParseNumbers(strings.Fields(elf_input))
		sum := 0
		for _, calories := range nums {
			sum += calories
		}

		elf_totals = append(elf_totals, sum)
		if sum > highest_elf {
			highest_elf = sum
		}
	}

	sort.Slice(elf_totals, func(i, j int) bool {
		return elf_totals[i] > elf_totals[j]

	})
	total := elf_totals[0] + elf_totals[1] + elf_totals[2]

	fmt.Println(fmt.Sprintf("Output Part 1: %d", highest_elf))
	fmt.Println(fmt.Sprintf("Output Part 2: %d", total))

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
