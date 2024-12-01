package main

import (
	"embed"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

//go:embed *.txt
var content embed.FS

func part1(input string) int {
	var left, right []int
	for _, line := range strings.Split(input, "\n") {
		numbers := strings.Split(line, " ")
		if len(numbers) > 1 {
			leftNumber, err := strconv.Atoi(numbers[0])
			if err != nil {
				log.Fatalf("Failed to parse left number \"%s\".\n", numbers[0])
			}
			left = append(left, leftNumber)

			rightNumber, err := strconv.Atoi(numbers[len(numbers)-1])
			if err != nil {
				log.Fatalf("Failed to parse right number \"%s\".\n", numbers[len(numbers)-1])
			}
			right = append(right, rightNumber)
		}
	}

	slices.Sort(left)
	slices.Sort(right)

	totalDistance := 0
	for i := 0; i < len(left); i++ {
		distance := left[i] - right[i]
		totalDistance += max(distance, -distance)
	}
	return totalDistance
}

func part2(input string) int {
	var left, right []int
	for _, line := range strings.Split(input, "\n") {
		numbers := strings.Split(line, " ")
		if len(numbers) > 1 {
			leftNumber, err := strconv.Atoi(numbers[0])
			if err != nil {
				log.Fatalf("Failed to parse left number \"%s\".\n", numbers[0])
			}
			left = append(left, leftNumber)

			rightNumber, err := strconv.Atoi(numbers[len(numbers)-1])
			if err != nil {
				log.Fatalf("Failed to parse right number \"%s\".\n", numbers[len(numbers)-1])
			}
			right = append(right, rightNumber)
		}
	}

	similarityScores := make(map[int]int)
	for _, l := range left {
		_, ok := similarityScores[l]
		if ok {
			continue
		}

		similarityScores[l] = 0
		for _, r := range right {
			if l == r {
				similarityScores[l]++
			}
		}
		similarityScores[l] *= l
	}

	totalSimilarityScore := 0
	for _, l := range left {
		totalSimilarityScore += similarityScores[l]
	}
	return totalSimilarityScore
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
