package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func checkUnique(s *string) bool {
	var a [256]bool
	for _, ascii := range *s {
		if a[ascii] {
			return false
		}
		a[ascii] = true
	}
	return true
}

func main() {
	fmt.Println("--- Running Day 6 ---")

	start := time.Now()

	file, err := os.ReadFile("./06.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Unable to read file")
		os.Exit(-1)
	}
	fileContent := string(file)

	raw_input := strings.Split(fileContent, "\n")
	line := raw_input[0]

	answer := 0
	for i := 0; i < len(line)-4; i++ {
		marker := ""
		for j := 0; j < 4; j++ {
			marker += string(line[i+j])
		}

		if checkUnique(&marker) {
			answer = i + 4
			break
		}
	}

	fmt.Printf("Output Part 1: %d\n", answer)

	answer = 0
	for i := 0; i < len(line)-14; i++ {
		marker := ""
		for j := 0; j < 14; j++ {
			marker += string(line[i+j])
		}

		if checkUnique(&marker) {
			answer = i + 14
			break
		}
	}

	fmt.Printf("Output Part 2: %d\n", answer)

	elapsed := time.Since(start)
	fmt.Printf("\nExecution took %s\n", elapsed)
}
