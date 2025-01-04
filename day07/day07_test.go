package main

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		filename string
		expected int
	}{
		{filename: "example.txt", expected: 3749},
		{filename: "input.txt", expected: 267566105056},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			input, err := os.ReadFile(tt.filename)
			if err != nil {
				t.Errorf("Failed to read file %v.", tt.filename)
			}
			actual := part1(string(input))
			if actual != tt.expected {
				t.Errorf("Expected %v. Received %v.", tt.expected, actual)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		filename string
		expected int
	}{
		{filename: "example.txt", expected: 11387},
		{filename: "input.txt", expected: 116094961956019},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			input, err := os.ReadFile(tt.filename)
			if err != nil {
				t.Errorf("Failed to read file %v.", tt.filename)
			}
			actual := part2(string(input))
			if actual != tt.expected {
				t.Errorf("Expected %v. Received %v.", tt.expected, actual)
			}
		})
	}
}
