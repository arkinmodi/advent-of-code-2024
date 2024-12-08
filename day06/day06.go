package main

import (
	"embed"
	"fmt"
	"log"
	"strings"
)

//go:embed *.txt
var content embed.FS

func part1(input string) int {
	var areaMap [][]rune
	for _, row := range strings.Split(input[:len(input)-1], "\n") {
		areaMap = append(areaMap, []rune(row))
	}

	type point struct {
		row int
		col int
	}

	var curr point
	for r := range len(areaMap) {
		for c := range len(areaMap[r]) {
			if areaMap[r][c] == '^' {
				curr = point{row: r, col: c}
				break
			}
		}
	}

	delta := point{row: -1, col: 0}
	next := point{}
	visited := make(map[point]struct{})
	inBounds := true
	for inBounds {
		visited[curr] = struct{}{}

		next.row = curr.row + delta.row
		next.col = curr.col + delta.col

		if 0 <= next.row &&
			next.row < len(areaMap) &&
			0 <= next.col &&
			next.col < len(areaMap[0]) {
			if areaMap[next.row][next.col] == '#' {
				if delta.row == -1 && delta.col == 0 {
					delta.row = 0
					delta.col = 1
				} else if delta.row == 0 && delta.col == 1 {
					delta.row = 1
					delta.col = 0
				} else if delta.row == 1 && delta.col == 0 {
					delta.row = 0
					delta.col = -1
				} else {
					delta.row = -1
					delta.col = 0
				}
				continue
			}
			curr.row = next.row
			curr.col = next.col
		} else {
			inBounds = false
		}
	}
	return len(visited)
}

func part2(input string) int {
	var areaMap [][]rune
	for _, row := range strings.Split(input[:len(input)-1], "\n") {
		areaMap = append(areaMap, []rune(row))
	}

	type point struct {
		row      int
		col      int
		deltaRow int
		deltaCol int
	}

	var curr point
	for r := range len(areaMap) {
		for c := range len(areaMap[r]) {
			if areaMap[r][c] == '^' {
				curr = point{row: r, col: c, deltaRow: -1, deltaCol: 0}
				break
			}
		}
	}

	run := func(modifiedAreaMap [][]rune, curr point) (map[point]struct{}, bool) {
		visited := make(map[point]struct{})
		next := point{}
		for {
			if _, ok := visited[curr]; ok {
				return visited, true
			}
			visited[curr] = struct{}{}

			next.row = curr.row + curr.deltaRow
			next.col = curr.col + curr.deltaCol

			if 0 <= next.row &&
				next.row < len(areaMap) &&
				0 <= next.col &&
				next.col < len(areaMap[0]) {
				if areaMap[next.row][next.col] == '#' {
					if curr.deltaRow == -1 && curr.deltaCol == 0 {
						curr.deltaRow = 0
						curr.deltaCol = 1
					} else if curr.deltaRow == 0 && curr.deltaCol == 1 {
						curr.deltaRow = 1
						curr.deltaCol = 0
					} else if curr.deltaRow == 1 && curr.deltaCol == 0 {
						curr.deltaRow = 0
						curr.deltaCol = -1
					} else {
						curr.deltaRow = -1
						curr.deltaCol = 0
					}
					continue
				}
				curr.row = next.row
				curr.col = next.col
			} else {
				break
			}
		}
		return visited, false
	}

	obstructionPositions, _ := run(areaMap, curr)
	result := 0
	visited := make(map[point]struct{})
	visited[point{row: curr.row, col: curr.col}] = struct{}{}
	for obstruction := range obstructionPositions {
		if _, ok := visited[point{row: obstruction.row, col: obstruction.col}]; ok {
			continue
		}
		visited[point{row: obstruction.row, col: obstruction.col}] = struct{}{}

		prev := areaMap[obstruction.row][obstruction.col]
		areaMap[obstruction.row][obstruction.col] = '#'
		if _, isLoop := run(areaMap, curr); isLoop {
			result++
		}
		if prev != '.' {
			fmt.Println(string(prev))
		}
		areaMap[obstruction.row][obstruction.col] = prev
	}
	return result
}

func main() {
	example, err := content.ReadFile("example.txt")
	if err != nil {
		log.Fatalf("Failed to read file \"example.txt\". Error: %v\n", err)
	}

	input, err := content.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read file \"input.txt\". Error: %v", err)
	}

	fmt.Printf("Part 1 Example:\t%v\n", part1(string(example)))
	fmt.Printf("Part 1:\t\t%v\n", part1(string(input)))
	fmt.Printf("Part 2 Example:\t%v\n", part2(string(example)))
	fmt.Printf("Part 2:\t\t%v\n", part2(string(input)))
}
