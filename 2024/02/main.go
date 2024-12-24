package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"slices"
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

// Check if the values in a report are within acceptable range.
func acceptableRange(reportString []string) (acceptableRange bool) {
	report := sliceStrToInt(reportString)
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
func safeReports(reports [][]string, dampener bool) (safe int) {
	for _, report := range reports {
		if acceptableRange(report) {
			safe++
		}
		// if the dampener is enabled, iterate through the report and see if removing a level would make the report safe.
		if !acceptableRange(report) && dampener {
			for i := range report {
				modifiedReport := slices.Concat(report[:i], report[i+1:])
				if acceptableRange(modifiedReport) {
					safe++
					break
				}
			}
		}
	}
	return
}

func main() {
	fmt.Println("Safe reports:", safeReports(parseFile(), false))
	fmt.Println("--Problem Dampener enabled--")
	fmt.Println("Safe reports:", safeReports(parseFile(), true))
}
