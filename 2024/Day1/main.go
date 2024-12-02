package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	start := time.Now()

	file, err := os.Open("./input.txt")

	check(err)

	defer file.Close()

	var col1, col2 []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		cols := strings.Fields(line)

		if len(cols) != 2 {
			fmt.Println("Skipping invalid line", line)
			continue
		}

		num1, err1 := strconv.Atoi(cols[0])
		num2, err2 := strconv.Atoi(cols[1])

		if err1 != nil || err2 != nil {
			fmt.Println("Skipping line due to conversion error:", line)
			continue
		}

		col1 = append(col1, num1)
		col2 = append(col2, num2)

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error reading file:", err)
		return
	}

	fmt.Println("Column 1: %d values loaded. \n", len(col1))
	fmt.Println("Column 2: %d values loaded. \n", len(col2))

	sort.Ints(col1)
	sort.Ints(col2)

	total := 0
	for i := 0; i < len(col1); i++ {
		tmp := col1[i] - col2[i]
		val := int(math.Abs(float64(tmp)))
		total += val
	}

	fmt.Println(total)
	elapsed := time.Since(start)
	fmt.Printf("Function took %s to execute \n", elapsed)

	// part 2

	counts := make(map[int]int)
	for _, num := range col1 {
		counts[num]++
	}

	similarityScore := 0

	for _, num := range col2 {
		similarityScore += num * counts[num]
	}

	fmt.Println("Similarity Score:", similarityScore)

}
