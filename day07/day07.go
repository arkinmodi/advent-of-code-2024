package main

import (
	"embed"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

//go:embed *.txt
var content embed.FS

func divmod(numerator, denominator int) (int, int) {
	quotient := numerator / denominator
	remainder := numerator % denominator
	return quotient, remainder
}

func part1IsEquationPossible(target int, nums []int) bool {
	if len(nums) == 0 {
		return false
	} else if len(nums) == 1 {
		return nums[0] == target
	}
	q, r := divmod(target, nums[len(nums)-1])
	if r == 0 && part1IsEquationPossible(q, nums[:len(nums)-1]) {
		return true
	}
	return part1IsEquationPossible(target-nums[len(nums)-1], nums[:len(nums)-1])
}

func part2IsEquationPossible(target int, nums []int) bool {
	if len(nums) == 0 {
		return false
	} else if len(nums) == 1 {
		return nums[0] == target
	}

	numDigits := func(n int) int {
		return len(strconv.Itoa(n))
	}

	endsWith := func(a, b int) bool {
		return strings.HasSuffix(strconv.Itoa(a), strconv.Itoa(b))
	}

	n := nums[len(nums)-1]
	q, r := divmod(target, n)
	if r == 0 && part2IsEquationPossible(q, nums[:len(nums)-1]) {
		return true
	} else if endsWith(target, n) && part2IsEquationPossible(target/(int(math.Pow10(numDigits(n)))), nums[:len(nums)-1]) {
		return true
	}
	return part2IsEquationPossible(target-n, nums[:len(nums)-1])
}

func part1(input string) int {
	type calibration struct {
		target int
		nums   []int
	}

	var calibrations []calibration
	for _, line := range strings.Split(input[:len(input)-1], "\n") {
		splitColon := strings.SplitN(line, ": ", 2)

		target, err := strconv.Atoi(splitColon[0])
		if err != nil {
			log.Fatalf("Failed to parse target number \"%s\".\n", splitColon[0])
		}

		var nums []int
		for _, str := range strings.Split(splitColon[1], " ") {
			n, err := strconv.Atoi(str)
			if err != nil {
				log.Fatalf("Failed to parse number \"%s\".\n", str)
			}
			nums = append(nums, n)
		}
		calibrations = append(calibrations, calibration{target, nums})
	}

	total := 0
	for _, cal := range calibrations {
		if part1IsEquationPossible(cal.target, cal.nums) {
			total += cal.target
		}
	}
	return total
}

func part2(input string) int {
	type calibration struct {
		target int
		nums   []int
	}

	var calibrations []calibration
	for _, line := range strings.Split(input[:len(input)-1], "\n") {
		splitColon := strings.SplitN(line, ": ", 2)

		target, err := strconv.Atoi(splitColon[0])
		if err != nil {
			log.Fatalf("Failed to parse target number \"%s\".\n", splitColon[0])
		}

		var nums []int
		for _, str := range strings.Split(splitColon[1], " ") {
			n, err := strconv.Atoi(str)
			if err != nil {
				log.Fatalf("Failed to parse number \"%s\".\n", str)
			}
			nums = append(nums, n)
		}
		calibrations = append(calibrations, calibration{target, nums})
	}

	total := 0
	for _, cal := range calibrations {
		if part2IsEquationPossible(cal.target, cal.nums) {
			total += cal.target
		}
	}
	return total
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
