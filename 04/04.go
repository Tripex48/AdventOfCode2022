package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
)

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

		p1 := mapset.NewSet[int]()
		for i := p1_nums[0]; i <= p1_nums[1]; i++ {
			p1.Add(i)
		}

		p2 := mapset.NewSet[int]()
		for i := p2_nums[0]; i <= p2_nums[1]; i++ {
			p2.Add(i)
		}

		if p1.IsSubset(p2) || p2.IsSubset(p1) {
			count_p1 += 1
		}

		if len(p1.Intersect(p2).ToSlice()) > 0 {
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
