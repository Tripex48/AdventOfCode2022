package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Pramod-Devireddy/go-exprtk"
)

type Monkey struct {
	nums []float64
	calc exprtk.GoExprtk
	div  int
	pos  int
	neg  int
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
				calc := monkeysTemp[input].calc
				calc.SetDoubleVariableValue("old", float64(item))
				val := calc.GetEvaluatedValue()

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

		calc := exprtk.NewExprtk()
		calc.SetExpression(strings.Split(lines[2], "=")[1])
		calc.AddDoubleVariable("old")
		calc.CompileExpression()

		div, _ := strconv.Atoi(strings.Fields(lines[3])[3])
		pos, _ := strconv.Atoi(strings.Fields(lines[4])[5])
		neg, _ := strconv.Atoi(strings.Fields(lines[5])[5])

		newMonkey := Monkey{
			nums: nums,
			calc: calc,
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
