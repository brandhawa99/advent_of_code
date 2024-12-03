package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// isSafeSequence checks if a sequence is safe
func isSafeSequence(levels []int) bool {
	// Check increasing sequence
	isIncreasing := true
	for i := 0; i < len(levels)-1; i++ {
		diff := levels[i+1] - levels[i]
		if diff < 1 || diff > 3 {
			isIncreasing = false
			break
		}
	}
	if isIncreasing {
		return true
	}

	// Check decreasing sequence
	isDecreasing := true
	for i := 0; i < len(levels)-1; i++ {
		diff := levels[i] - levels[i+1]
		if diff < 1 || diff > 3 {
			isDecreasing = false
			break
		}
	}
	return isDecreasing
}

// isSafeSequenceWithDampener checks if sequence is safe with one allowed error
func isSafeSequenceWithDampener(levels []int) bool {
	// Check original sequence
	if isSafeSequence(levels) {
		return true
	}

	// Try removing each level
	for i := 0; i < len(levels); i++ {
		// Create reduced sequence by removing one level
		reducedLevels := make([]int, 0, len(levels)-1)
		reducedLevels = append(reducedLevels, levels[:i]...)
		reducedLevels = append(reducedLevels, levels[i+1:]...)

		// Check if reduced sequence is safe
		if isSafeSequence(reducedLevels) {
			return true
		}
	}

	return false
}

func main() {
	// Read input from file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read reports
	var reports []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		reports = append(reports, scanner.Text())
	}

	// Part One
	partOneSolution := 0
	for _, report := range reports {
		levels := parseReport(report)
		if isSafeSequence(levels) {
			partOneSolution++
		}
	}
	fmt.Println("Part One - Safe Reports:", partOneSolution)

	// Part Two
	partTwoSolution := 0
	for _, report := range reports {
		levels := parseReport(report)
		if isSafeSequenceWithDampener(levels) {
			partTwoSolution++
		}
	}
	fmt.Println("Part Two - Safe Reports:", partTwoSolution)
}

// parseReport converts a string of space-separated numbers to []int
func parseReport(report string) []int {
	parts := strings.Fields(report)
	levels := make([]int, len(parts))
	for i, part := range parts {
		levels[i], _ = strconv.Atoi(part)
	}
	return levels
}
