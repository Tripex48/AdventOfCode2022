package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"
)

func getIntersection(setA map[rune]struct{}, setB map[rune]struct{}) map[rune]struct{} {
	newSet := make(map[rune]struct{})

	for chr := range setA {
		if _, ok := setB[chr]; ok {
			newSet[chr] = struct{}{}
		}
	}

	return newSet
}

func main() {
	fmt.Println("--- Running Day 3 ---")

	start := time.Now()

	file, err := os.ReadFile("./03.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Unable to read file")
		os.Exit(-1)
	}
	fileContent := string(file)

	input := strings.Split(strings.TrimSpace(fileContent), "\n")

	priority := 0
	for _, line := range input {
		cmp1 := make(map[rune]struct{})
		for _, chr := range line[:len(line)/2] {
			cmp1[chr] = struct{}{}
		}

		cmp2 := make(map[rune]struct{})
		for _, chr := range line[len(line)/2:] {
			cmp2[chr] = struct{}{}
		}

		dup := getIntersection(cmp1, cmp2)
		keys := make([]rune, 0, len(dup))
		for k := range dup {
			keys = append(keys, k)
			break
		}

		if unicode.IsLower(keys[0]) {
			priority += int(keys[0]) - 96
		} else {
			priority += int(keys[0]) - 64 + 26
		}
	}

	fmt.Printf("Output Part 1: %d\n", priority)

	priority = 0
	for i := 0; i < len(input); i += 3 {
		cmp1 := make(map[rune]struct{})
		for _, chr := range input[i] {
			cmp1[chr] = struct{}{}
		}

		cmp2 := make(map[rune]struct{})
		for _, chr := range input[i+1] {
			cmp2[chr] = struct{}{}
		}

		cmp3 := make(map[rune]struct{})
		for _, chr := range input[i+2] {
			cmp3[chr] = struct{}{}
		}

		dup := getIntersection(cmp1, cmp2)
		dup2 := getIntersection(dup, cmp3)

		keys := make([]rune, 0, len(dup2))
		for k := range dup2 {
			keys = append(keys, k)
			break
		}

		if unicode.IsLower(keys[0]) {
			priority += int(keys[0]) - 96
		} else {
			priority += int(keys[0]) - 64 + 26
		}
	}

	fmt.Printf("Output Part 2: %d\n", priority)

	elapsed := time.Since(start)
	fmt.Printf("\nExecution took %s\n", elapsed)
}
