package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("--- Running Day 10 ---")

	start := time.Now()

	file, err := os.ReadFile("./10.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Unable to read file")
		os.Exit(-1)
	}
	fileContent := string(file)

	lines := strings.Split(strings.TrimSpace(fileContent), "\n")

	x := 1
	stack := []int{}

	for _, line := range lines {
		if strings.Contains(line, "noop") {
			stack = append(stack, x)
		} else {
			val, _ := strconv.Atoi(strings.Split(line, " ")[1])
			stack = append(stack, x)
			stack = append(stack, x)
			x += val
		}
	}

	sum := 0
	sums := []int{20, 60, 100, 140, 180, 220}

	for _, i := range sums {
		sum += i * stack[i-1]
	}

	fmt.Printf("Output Part 1: %d\n", sum)

	str := "\n"
	for i := 0; i < len(stack); i += 40 {
		for j := 0; j < 40; j++ {
			if math.Abs(float64(stack[i+j]-j)) <= 1 {
				str += "#"
			} else {
				str += " "
			}
		}
		str += "\n"
	}

	fmt.Printf("Output Part 2: %s\n", str)

	elapsed := time.Since(start)
	fmt.Printf("\nExecution took %s\n", elapsed)
}
