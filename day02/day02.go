package main

import (
	"embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed *.txt
var content embed.FS

func part1(input string) int {
	safeReports := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		var levels []int
		for _, l := range strings.Fields(line) {
			number, err := strconv.Atoi(l)
			if err != nil {
				log.Fatalf("Failed to parse \"%v\" to an integer.\n", l)
			}
			levels = append(levels, number)
		}

		isDecreasing := levels[1] < levels[0]
		isSafeReport := true
		for i := range len(levels) - 1 {
			diff := levels[i+1] - levels[i]
			if diff < 0 {
				diff *= -1
			}

			if !(1 <= diff && diff <= 3 && levels[i+1] < levels[i] == isDecreasing) {
				isSafeReport = false
				break
			}
		}

		if isSafeReport {
			safeReports++
		}
	}
	return safeReports
}

func part2(input string) int {
	safeReports := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		var levels []int
		for _, l := range strings.Fields(line) {
			number, err := strconv.Atoi(l)
			if err != nil {
				log.Fatalf("Failed to parse \"%v\" to an integer.\n", l)
			}
			levels = append(levels, number)
		}

		var possibleReports [][]int
		possibleReports = append(possibleReports, levels)
		for i := range len(levels) {
			newReport := make([]int, len(levels[:i]))
			copy(newReport, levels[:i])
			newReport = append(newReport, levels[i+1:]...)
			possibleReports = append(possibleReports, newReport)
		}

		isSafeReport := true
		for _, report := range possibleReports {
			isSafeReport = true
			isDecreasing := report[1] < report[0]
			for i := range len(report) - 1 {
				diff := report[i+1] - report[i]
				if diff < 0 {
					diff *= -1
				}

				if !(1 <= diff && diff <= 3 && report[i+1] < report[i] == isDecreasing) {
					isSafeReport = false
					break
				}
			}

			if isSafeReport {
				break
			}
		}

		if isSafeReport {
			safeReports++
		}
	}
	return safeReports
}

func main() {
	example, err := content.ReadFile("example.txt")
	if err != nil {
		log.Fatalf("Failed to read file \"example.txt\". Error: %v\n", err)
	}

	input, err := content.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read file \"intput.txt\". Error: %v\n", err)
	}

	fmt.Printf("Part 1 Example:\t%v\n", part1(string(example)))
	fmt.Printf("Part 1:\t\t%v\n", part1(string(input)))
	fmt.Printf("Part 2 Example:\t%v\n", part2(string(example)))
	fmt.Printf("Part 2:\t\t%v\n", part2(string(input)))
}
