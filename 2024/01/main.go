package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

// function parseAndSort() reads the file and returns two slices of integers
func parseAndSort() ([]int, []int) {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	col1, col2 := []int{}, []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var num1, num2 int
		fmt.Sscanf(line, "%d   %d", &num1, &num2)
		col1 = append(col1, num1)
		col2 = append(col2, num2)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	slices.Sort(col1)
	slices.Sort(col2)
	return col1, col2
}

// function positive() converts a given int (num) to positive.
func positive(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// function totalDiff() compares each value of two arrays and returns the total difference
func totalDiff(col1, col2 []int) int {
	totalDiff := 0
	for i, v := range col1 {
		totalDiff += positive(v - col2[i])
	}
	return totalDiff
}

func main() {
	fmt.Println("Total difference between columns:", totalDiff(parseAndSort()))
}
