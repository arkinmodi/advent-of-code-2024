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
	inputRows := strings.Split(strings.TrimSpace(input), "\n")
	pageOrderingRules := make(map[int][]int)
	i := 0
	for {
		if i >= len(inputRows) || inputRows[i] == "" {
			i++
			break
		}

		rule := strings.Split(inputRows[i], "|")
		key, err := strconv.Atoi(rule[0])
		if err != nil {
			log.Fatalf("Failed to parse \"%v\" into an integer.\n", rule[0])
		}

		value, err := strconv.Atoi(rule[1])
		if err != nil {
			log.Fatalf("Failed to parse \"%v\" into an integer.\n", rule[1])
		}

		pageOrderingRules[key] = append(pageOrderingRules[key], value)
		i++
	}

	var updates [][]int
	for {
		if i >= len(inputRows) {
			break
		}
		var row []int
		for _, page := range strings.Split(inputRows[i], ",") {
			val, err := strconv.Atoi(page)
			if err != nil {
				log.Fatalf("Failed to parse %q as an integer.\n", page)
			}
			row = append(row, val)
		}
		updates = append(updates, row)
		i++
	}

	middlePageSum := 0
	for _, update := range updates {
		isSorted := true
		for i, page := range update {
			for j := range len(update) - i {
				rule, ok := pageOrderingRules[update[i+j]]
				if !ok {
					continue
				}
				if slices.Contains(rule, page) {
					isSorted = false
					break
				}
			}

			if !isSorted {
				break
			}
		}

		if isSorted {
			middlePageSum += update[len(update)/2]
		}
	}
	return middlePageSum
}

func part2(input string) int {
	inputRows := strings.Split(strings.TrimSpace(input), "\n")
	pageOrderingRules := make(map[int][]int)
	i := 0
	for {
		if i >= len(inputRows) || inputRows[i] == "" {
			i++
			break
		}

		rule := strings.Split(inputRows[i], "|")
		key, err := strconv.Atoi(rule[0])
		if err != nil {
			log.Fatalf("Failed to parse \"%v\" into an integer.\n", rule[0])
		}

		value, err := strconv.Atoi(rule[1])
		if err != nil {
			log.Fatalf("Failed to parse \"%v\" into an integer.\n", rule[1])
		}

		pageOrderingRules[key] = append(pageOrderingRules[key], value)
		i++
	}

	var updates [][]int
	for {
		if i >= len(inputRows) {
			break
		}

		var row []int
		for _, page := range strings.Split(inputRows[i], ",") {
			val, err := strconv.Atoi(page)
			if err != nil {
				log.Fatalf("Failed to parse %q as an integer.\n", page)
			}
			row = append(row, val)
		}
		updates = append(updates, row)
		i++
	}

	sliceCommon := func(a []int, b []int) []int {
		if len(a) == 0 || len(b) == 0 {
			return []int{}
		}

		var c []int
		for _, v := range a {
			if slices.Contains(b, v) {
				c = append(c, v)
			}
		}
		return c
	}

	middlePageSum := 0
	for _, update := range updates {
		isSorted := true
		for i, page := range update {
			for j := range len(update) - i {
				rule, ok := pageOrderingRules[update[i+j]]
				if !ok {
					continue
				}
				if slices.Contains(rule, page) {
					isSorted = false
					break
				}
			}

			if !isSorted {
				break
			}
		}

		if isSorted {
			continue
		}

		newUpdate := []int{}
		for len(update) > 0 {
			for i, page := range update {
				rule, ok := pageOrderingRules[page]
				if !ok || len(sliceCommon(rule, update)) == 0 {
					newUpdate = append(newUpdate, page)
					update = append(update[:i], update[i+1:]...)
					break
				}
			}
		}
		middlePageSum += newUpdate[len(newUpdate)/2]
	}
	return middlePageSum
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
