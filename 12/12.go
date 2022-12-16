package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Workiva/go-datastructures/queue"
)

type Coord struct {
	x int
	y int
}

type Entry struct {
	dist  int
	coord Coord
}

func (e Entry) Compare(other queue.Item) int {
	if e.dist < other.(Entry).dist {
		return -1
	} else if e.dist == other.(Entry).dist {
		return 0
	} else {
		return 1
	}
}

func GetNeighbours(x int, y int) []Coord {
	neighbours := []Coord{}
	neighbours = append(neighbours, Coord{x: x - 1, y: y})
	neighbours = append(neighbours, Coord{x: x, y: y - 1})
	neighbours = append(neighbours, Coord{x: x + 1, y: y})
	neighbours = append(neighbours, Coord{x: x, y: y + 1})
	return neighbours
}

func bfs(coords map[Coord]rune, start Coord, end Coord, part int) int {
	visited := make(map[Coord]struct{})
	visited[end] = struct{}{}
	queue := queue.NewPriorityQueue(0, false)
	queue.Put(Entry{dist: 0, coord: end})

	for queue.Len() > 0 {
		queue_item, _ := queue.Get(1)
		dist := queue_item[0].(Entry).dist
		x := queue_item[0].(Entry).coord.x
		y := queue_item[0].(Entry).coord.y

		val := coords[Coord{x: x, y: y}]
		for _, coord := range GetNeighbours(x, y) {
			if _, ok := coords[coord]; !ok {
				continue
			}
			if _, ok := visited[coord]; ok {
				continue
			}
			if int(coords[coord])-int(val) < -1 {
				continue
			}

			if part == 1 {
				if coord == start {
					return dist + 1
				}
			} else {
				if coords[coord] == 'a' {
					return dist + 1
				}
			}

			visited[coord] = struct{}{}
			queue.Put(Entry{dist: dist + 1, coord: coord})
		}
	}

	return 0
}

func main() {
	fmt.Println("--- Running Day 12 ---")

	start := time.Now()

	file, err := os.ReadFile("./12.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Unable to read file")
		os.Exit(-1)
	}
	fileContent := string(file)

	input := strings.Split(strings.TrimSpace(fileContent), "\n")

	coords := make(map[Coord]rune)
	startPoint := Coord{x: 0, y: 0}
	endPoint := Coord{x: -1, y: -1}
	for row, line := range input {
		for col, chr := range line {
			if chr == 'S' {
				startPoint = Coord{x: row, y: col}
				coords[startPoint] = 'a'
			} else if chr == 'E' {
				endPoint = Coord{x: row, y: col}
				coords[endPoint] = 'z'
			} else {
				coords[Coord{x: row, y: col}] = chr
			}
		}
	}

	fmt.Printf("Output Part 1: %d\n", bfs(coords, startPoint, endPoint, 1))
	fmt.Printf("Output Part 2: %d\n", bfs(coords, startPoint, endPoint, 2))

	elapsed := time.Since(start)
	fmt.Printf("\nExecution took %s\n", elapsed)
}
