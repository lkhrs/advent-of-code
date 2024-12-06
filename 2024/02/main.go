package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Reads the file and returns a slice of strings containing a series of numbers.
// Example line: 4 8 15 16 23 42 => "4", "8", "15", "16", "23", "42"
func parseFile() [][]string {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}

	r := csv.NewReader(file)
	r.Comma = ' '
	r.FieldsPerRecord = -1
	reports, err := r.ReadAll()
	if err != nil {
		log.Fatal("Error reading input file: ", err)
	}

	return reports
}

// Converts a given int (num) to positive.
func positive(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// Converts a slice of strings to a slice of integers.
func sliceStrToInt(str []string) []int {
	intSlice := []int{}
	for _, s := range str {
		num, err := strconv.Atoi(string(s))
		if err != nil {
			log.Fatal(err)
		}
		intSlice = append(intSlice, num)
	}
	return intSlice
}

// Check if a report is within acceptable range.
func acceptableRange(report []int) bool {
	if len(report) < 2 {
		return false
	}
	increasing := report[1] > report[0]
	for i, _ := range report {
		if i > 0 {
			if positive(report[i]-report[i-1]) > 3 {
				return false
			}
			if increasing != (report[i] > report[i-1]) {
				return false
			}
			if report[i] == report[i-1] {
				return false
			}
		}
	}
	return true
}

// Iterate through the reports and count the safe ones.
func safeReports(reports [][]string) int {
	safe := 0
	for _, report := range reports {
		if acceptableRange(sliceStrToInt(report)) {
			safe++
		}
	}
	return safe
}

func main() {
	fmt.Println("Safe reports:", safeReports(parseFile()))
}
