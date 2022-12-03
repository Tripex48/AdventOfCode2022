package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"

	mapset "github.com/deckarep/golang-set/v2"
)

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
		cmp1 := mapset.NewSet[rune]()
		for _, chr := range line[:len(line)/2] {
			cmp1.Add(chr)
		}

		cmp2 := mapset.NewSet[rune]()
		for _, chr := range line[len(line)/2:] {
			cmp2.Add(chr)
		}

		dup, _ := cmp1.Intersect(cmp2).Pop()
		if unicode.IsLower(dup) {
			priority += int(dup) - 96
		} else {
			priority += int(dup) - 64 + 26
		}
	}

	fmt.Println(fmt.Sprintf("Output Part 1: %d", priority))

	priority = 0
	for i := 0; i < len(input); i += 3 {
		cmp1 := mapset.NewSet[rune]()
		for _, chr := range input[i] {
			cmp1.Add(chr)
		}

		cmp2 := mapset.NewSet[rune]()
		for _, chr := range input[i+1] {
			cmp2.Add(chr)
		}

		cmp3 := mapset.NewSet[rune]()
		for _, chr := range input[i+2] {
			cmp3.Add(chr)
		}

		dup, _ := cmp1.Intersect(cmp2).Intersect(cmp3).Pop()
		if unicode.IsLower(dup) {
			priority += int(dup) - 96
		} else {
			priority += int(dup) - 64 + 26
		}

	}

	fmt.Println(fmt.Sprintf("Output Part 2: %d", priority))

	elapsed := time.Since(start)
	fmt.Printf("\nExecution took %s\n", elapsed)
}
