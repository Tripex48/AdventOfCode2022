package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func reverse(input []string) {
	inputLen := len(input)
	inputMid := inputLen / 2

	for i := 0; i < inputMid; i++ {
		j := inputLen - i - 1

		input[i], input[j] = input[j], input[i]
	}
}

func reverseString(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func main() {
	fmt.Println("--- Running Day 5 ---")

	start := time.Now()

	file, err := os.ReadFile("./05.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Unable to read file")
		os.Exit(-1)
	}
	fileContent := string(file)

	raw_input := strings.Split(fileContent, "\n\n")
	input := strings.Split(raw_input[0], "\n")
	reverse(input)
	moves := strings.Split(strings.TrimSpace(raw_input[1]), "\n")

	len_s := strings.Fields(input[0])
	reverse(len_s)
	length, _ := strconv.Atoi(len_s[0])

	stacks_p1 := make([]string, length)
	stacks_p2 := make([]string, length)
	for _, line := range input[1:] {
		index := 0
		for i := 1; i <= len(line); i += 4 {
			if string(line[i:i+1]) != "" {
				stacks_p1[index] += string(line[i : i+1])
				stacks_p2[index] += string(line[i : i+1])
				index += 1
			}
		}
	}

	for i := 0; i < length; i++ {
		stacks_p1[i] = strings.TrimSpace(stacks_p1[i])
		stacks_p2[i] = strings.TrimSpace(stacks_p2[i])
	}

	for _, line := range moves {

		raw_str := strings.Fields(line)
		num, _ := strconv.Atoi(raw_str[1])
		init, _ := strconv.Atoi(raw_str[3])
		dest, _ := strconv.Atoi(raw_str[5])

		letters := ""
		for i := 0; i < num; i++ {
			letter := stacks_p1[init-1][len(stacks_p1[init-1])-1]
			stacks_p1[dest-1] += string(letter)
			stacks_p1[init-1] = stacks_p1[init-1][:len(stacks_p1[init-1])-1]

			letter_p2 := stacks_p2[init-1][len(stacks_p2[init-1])-1]
			letters += string(letter_p2)
			stacks_p2[init-1] = stacks_p2[init-1][:len(stacks_p2[init-1])-1]
		}
		letters = reverseString(letters)
		stacks_p2[dest-1] += letters
	}

	answer := ""
	for i := 0; i < length; i++ {
		if len(stacks_p1[i]) > 0 {
			answer += string(stacks_p1[i][len(stacks_p1[i])-1])
		}
	}

	fmt.Println(fmt.Sprintf("Output Part 1: %s", answer))

	answer = ""
	for i := 0; i < length; i++ {
		if len(stacks_p2[i]) > 0 {
			answer += string(stacks_p2[i][len(stacks_p2[i])-1])
		}
	}

	fmt.Println(fmt.Sprintf("Output Part 2: %s", answer))

	elapsed := time.Since(start)
	fmt.Printf("\nExecution took %s\n", elapsed)
}
