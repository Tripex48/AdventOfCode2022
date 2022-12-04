package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func getIntersection(setA map[int]struct{}, setB map[int]struct{}) map[int]struct{} {
	newSet := make(map[int]struct{})

	for num := range setA {
		if _, ok := setB[num]; ok {
			newSet[num] = struct{}{}
		}
	}

	return newSet
}

func isSuperset(set map[int]struct{}, subset map[int]struct{}) bool {
	for elem := range subset {
		if _, ok := set[elem]; !ok {
			return false
		}
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

		p1 := make(map[int]struct{})
		for i := p1_nums[0]; i <= p1_nums[1]; i++ {
			p1[i] = struct{}{}
		}

		p2 := make(map[int]struct{})
		for i := p2_nums[0]; i <= p2_nums[1]; i++ {
			p2[i] = struct{}{}
		}

		if isSuperset(p1, p2) || isSuperset(p2, p1) {
			count_p1 += 1
		}

		dup := getIntersection(p1, p2)
		keys := make([]int, 0, len(dup))
		for k := range dup {
			keys = append(keys, k)
			break
		}

		if len(keys) > 0 {
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
