package main

import (
	"embed"
	"fmt"
	"log"
	"strings"
)

//go:embed *.txt
var content embed.FS

func printGrid(grid [][]rune) {
	var sb strings.Builder
	for _, r := range grid {
		sb.WriteString(string(r))
		sb.WriteString("\n")
	}
	fmt.Println(sb.String())
}

func part1(input string) int {
	var wordSearch [][]rune
	for _, row := range strings.Split(input[:len(input)-1], "\n") {
		wordSearch = append(wordSearch, []rune(row))
	}

	xmasCount := 0
	for r := range len(wordSearch) {
		for c := range len(wordSearch[r]) {
			if wordSearch[r][c] == 'X' {

				// Up
				if 0 <= r-3 &&
					wordSearch[r-1][c] == 'M' &&
					wordSearch[r-2][c] == 'A' &&
					wordSearch[r-3][c] == 'S' {
					xmasCount++
				}

				// Down
				if r+3 < len(wordSearch) &&
					wordSearch[r+1][c] == 'M' &&
					wordSearch[r+2][c] == 'A' &&
					wordSearch[r+3][c] == 'S' {
					xmasCount++
				}

				// Left
				if 0 <= c-3 &&
					wordSearch[r][c-1] == 'M' &&
					wordSearch[r][c-2] == 'A' &&
					wordSearch[r][c-3] == 'S' {
					xmasCount++
				}

				// Right
				if c+3 < len(wordSearch[r]) &&
					wordSearch[r][c+1] == 'M' &&
					wordSearch[r][c+2] == 'A' &&
					wordSearch[r][c+3] == 'S' {
					xmasCount++
				}

				// Up-Left
				if 0 <= r-3 &&
					0 <= c-3 &&
					wordSearch[r-1][c-1] == 'M' &&
					wordSearch[r-2][c-2] == 'A' &&
					wordSearch[r-3][c-3] == 'S' {
					xmasCount++
				}

				// Up-Right
				if 0 <= r-3 &&
					c+3 < len(wordSearch[r]) &&
					wordSearch[r-1][c+1] == 'M' &&
					wordSearch[r-2][c+2] == 'A' &&
					wordSearch[r-3][c+3] == 'S' {
					xmasCount++
				}

				// Down-Left
				if r+3 < len(wordSearch) &&
					0 <= c-3 &&
					wordSearch[r+1][c-1] == 'M' &&
					wordSearch[r+2][c-2] == 'A' &&
					wordSearch[r+3][c-3] == 'S' {
					xmasCount++
				}

				// Down-Right
				if r+3 < len(wordSearch) &&
					c+3 < len(wordSearch[r]) &&
					wordSearch[r+1][c+1] == 'M' &&
					wordSearch[r+2][c+2] == 'A' &&
					wordSearch[r+3][c+3] == 'S' {
					xmasCount++
				}
			}
		}
	}

	return xmasCount
}

func part2(input string) int {
	var wordSearch [][]rune
	for _, row := range strings.Split(input[:len(input)-1], "\n") {
		wordSearch = append(wordSearch, []rune(row))
	}

	// Bottom-Left to Top-Right
	isForwardDiagonal := func(r int, c int) bool {
		return (wordSearch[r+1][c-1] == 'M' && wordSearch[r-1][c+1] == 'S') ||
			(wordSearch[r+1][c-1] == 'S' && wordSearch[r-1][c+1] == 'M')
	}

	// Bottom-Right to Top-Left
	isBackwardDiagonal := func(r int, c int) bool {
		return (wordSearch[r-1][c-1] == 'M' && wordSearch[r+1][c+1] == 'S') ||
			(wordSearch[r-1][c-1] == 'S' && wordSearch[r+1][c+1] == 'M')
	}

	xmasCount := 0
	for r := range len(wordSearch) {
		for c := range len(wordSearch[r]) {
			if wordSearch[r][c] == 'A' {
				if !(0 <= r-1 &&
					r+1 < len(wordSearch) &&
					0 <= c-1 &&
					c+1 < len(wordSearch[r])) {
					continue
				}

				if isForwardDiagonal(r, c) && isBackwardDiagonal(r, c) {
					xmasCount++
				}
			}
		}
	}

	return xmasCount
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
