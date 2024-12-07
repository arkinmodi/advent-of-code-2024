package main

import (
	"embed"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

//go:embed *.txt
var content embed.FS

func part1(input string) int {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	result := 0
	for _, mul := range matches {
		left, err := strconv.Atoi(mul[1])
		if err != nil {
			log.Fatalf("Failed to parse %v into an integer.\n", mul[1])
		}

		right, err := strconv.Atoi(mul[2])
		if err != nil {
			log.Fatalf("Failed to parse %v into an integer.\n", mul[2])
		}

		result += left * right
	}
	return result
}

func part2(input string) int {
	re := regexp.MustCompile(`don't\(\)|do\(\)|mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	instructionsEnabled := true
	result := 0
	for _, instruction := range matches {
		if instruction[0] == "do()" {
			instructionsEnabled = true
			continue
		} else if instruction[0] == "don't()" {
			instructionsEnabled = false
			continue
		}

		if !instructionsEnabled {
			continue
		}

		left, err := strconv.Atoi(instruction[1])
		if err != nil {
			log.Fatalf("Failed to parse %v into an integer.\n", instruction[1])
		}

		right, err := strconv.Atoi(instruction[2])
		if err != nil {
			log.Fatalf("Failed to parse %v into an integer.\n", instruction[2])
		}

		result += left * right
	}
	return result
}

func main() {
	example_part1, err := content.ReadFile("example_part1.txt")
	if err != nil {
		log.Fatalf("Failed to read file \"example.txt\". Error: %v\n", err)
	}

	example_part2, err := content.ReadFile("example_part2.txt")
	if err != nil {
		log.Fatalf("Failed to read file \"example.txt\". Error: %v\n", err)
	}

	input, err := content.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read file \"input.txt\". Error: %v", err)
	}

	fmt.Printf("Part 1 Example:\t%v\n", part1(string(example_part1)))
	fmt.Printf("Part 1:\t\t%v\n", part1(string(input)))
	fmt.Printf("Part 2 Example:\t%v\n", part2(string(example_part2)))
	fmt.Printf("Part 2:\t\t%v\n", part2(string(input)))
}
