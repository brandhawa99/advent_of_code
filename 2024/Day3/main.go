package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// Input string with corrupted memory
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input string

	for scanner.Scan() {
		input += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	// Regular expression to find valid mul(X,Y) patterns
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	// Find all matches
	matches := re.FindAllStringSubmatch(input, -1)

	// Initialize sum
	totalSum := 0

	// Process each valid match
	for _, match := range matches {
		// Extract numbers from the match
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])

		// Compute product and add to total sum
		totalSum += num1 * num2
	}

	// Print the result
	fmt.Println("Total Sum:", totalSum)

	// Regular expressions
	mulRe := regexp.MustCompile(`mul\((\d+),(\d+)\)`) // Matches mul(X,Y)
	doRe := regexp.MustCompile(`do\(\)`)              // Matches do()
	dontRe := regexp.MustCompile(`don't\(\)`)         // Matches don't()

	// Track whether mul instructions are enabled
	mulEnabled := true
	totalSum = 0

	// Sequentially process instructions in the input
	for _, token := range regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`).FindAllString(input, -1) {
		if doRe.MatchString(token) {
			// Enable mul instructions
			mulEnabled = true
		} else if dontRe.MatchString(token) {
			// Disable mul instructions
			mulEnabled = false
		} else if mulRe.MatchString(token) && mulEnabled {
			// Extract and process mul(X,Y)
			matches := mulRe.FindStringSubmatch(token)
			num1, _ := strconv.Atoi(matches[1])
			num2, _ := strconv.Atoi(matches[2])
			totalSum += num1 * num2
		}
	}

	// Print the result
	fmt.Println("Total Sum:", totalSum)

}
