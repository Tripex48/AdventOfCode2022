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
	fmt.Println("--- Running Day 7 ---")

	start := time.Now()

	file, err := os.ReadFile("./07.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Unable to read file")
		os.Exit(-1)
	}
	fileContent := string(file)

	lines := strings.Split(strings.TrimSpace(fileContent), "\n")

	directories := make(map[string]int)
	path := []string{}

	for _, line := range lines {
		words := strings.Fields(strings.TrimSpace(line))
		if words[1] == "cd" {
			if words[2] == ".." {
				path = path[:len(path)-1]
			} else {
				path = append(path, words[2])
			}
		} else if words[1] == "ls" {
			continue
		} else {
			if words[0] != "dir" {
				size, _ := strconv.Atoi(words[0])
				for i := 0; i < len(path)+1; i++ {
					newPath := strings.ReplaceAll(strings.Join(path[:i], "/"), "//", "/")
					if newPath != "" {
						directories[newPath] += size
					}
				}
			}
		}
	}

	answer_p1 := 0
	for _, value := range directories {
		if value <= 100000 {
			answer_p1 += value
		}
	}

	fmt.Printf("Output Part 1: %d\n", answer_p1)

	usedSpace := directories["/"]
	freeSpace := 70000000 - usedSpace

	keys := make([]string, 0, len(directories))
	for key := range directories {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return directories[keys[i]] > directories[keys[j]]
	})

	answer_p2 := 0
	for _, k := range keys {
		if freeSpace+directories[k] >= 30000000 {
			answer_p2 = directories[k]
		}
	}

	fmt.Printf("Output Part 2: %d\n", answer_p2)

	elapsed := time.Since(start)
	fmt.Printf("\nExecution took %s\n", elapsed)
}
