package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Monkey struct {
	nums []float64
	args []int
	op   string
	div  int
	pos  int
	neg  int
}

func eval(num1 float64, op string, num2 float64, item float64) float64 {
	if num1 == -1 {
		num1 = item
	}
	if num2 == -1 {
		num2 = item
	}

	x := float64(0)
	switch op {
	case "+":
		x = num1 + num2
		break
	case "-":
		x = num1 - num2
		break
	case "*":
		x = num1 * num2
		break
	}

	return x
}

func solve(monkeys []Monkey, iter int, part int) int {
	counts := []int{}
	for i := 0; i < len(monkeys); i++ {
		counts = append(counts, 0)
	}

	monkeysTemp := make([]Monkey, len(monkeys))
	copy(monkeysTemp, monkeys)

	mod := float64(1)
	if part == 2 {
		for _, input := range monkeysTemp {
			mod *= float64(input.div)
		}
	}

	for i := 0; i < iter; i++ {
		for input := 0; input < len(monkeysTemp); input++ {
			for _, item := range monkeysTemp[input].nums {
				val := eval(float64(monkeysTemp[input].args[0]), monkeysTemp[input].op, float64(monkeysTemp[input].args[1]), item)

				if part == 1 {
					val = math.Floor(val / 3)
				} else {
					val = math.Mod(val, mod)
				}

				if math.Mod(val, float64(monkeysTemp[input].div)) == 0 {
					monkeysTemp[monkeysTemp[input].pos].nums = append(monkeysTemp[monkeysTemp[input].pos].nums, val)
				} else {
					monkeysTemp[monkeysTemp[input].neg].nums = append(monkeysTemp[monkeysTemp[input].neg].nums, val)
				}
			}

			counts[input] += len(monkeysTemp[input].nums)
			monkeysTemp[input].nums = []float64{}
		}
	}

	sort.Ints(counts)
	return counts[len(counts)-1] * counts[len(counts)-2]
}

func main() {
	fmt.Println("--- Running Day 11 ---")

	start := time.Now()

	file, err := os.ReadFile("./11.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Unable to read file")
		os.Exit(-1)
	}
	fileContent := string(file)

	input := strings.Split(strings.TrimSpace(fileContent), "\n\n")

	monkeys := []Monkey{}

	for _, monkey := range input {
		lines := strings.Split(monkey, "\n")

		nums := []float64{}
		for _, x := range strings.Split(strings.Split(lines[1], ": ")[1], ", ") {
			num, _ := strconv.Atoi(x)
			nums = append(nums, float64(num))
		}

		formula := strings.TrimSpace(strings.Split(lines[2], "=")[1])
		calc := strings.Fields(formula)
		args := []int{-1, -1}
		op := calc[1]
		if !strings.Contains(calc[0], "old") {
			num, _ := strconv.Atoi(calc[0])
			args[0] = num
		}
		if !strings.Contains(calc[2], "old") {
			num, _ := strconv.Atoi(calc[2])
			args[1] = num
		}

		div, _ := strconv.Atoi(strings.Fields(lines[3])[3])
		pos, _ := strconv.Atoi(strings.Fields(lines[4])[5])
		neg, _ := strconv.Atoi(strings.Fields(lines[5])[5])

		newMonkey := Monkey{
			nums: nums,
			args: args,
			op:   op,
			div:  div,
			pos:  pos,
			neg:  neg,
		}

		monkeys = append(monkeys, newMonkey)
	}

	fmt.Printf("Output Part 1: %d\n", solve(monkeys, 20, 1))

	fmt.Printf("Output Part 2: %d\n", solve(monkeys, 10000, 2))

	elapsed := time.Since(start)
	fmt.Printf("\nExecution took %s\n", elapsed)
}
