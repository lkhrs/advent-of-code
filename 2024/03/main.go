package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func stringToI(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	data := ""
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		data += scanner.Text()
	}
	
	// Match valid operations
	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	operations := regex.FindAllStringSubmatch(data, -1)
	
	// Compute the instructions
	total := 0
	for _, num := range operations {
		total += stringToI(num[1]) * stringToI(num[2])
	}
	fmt.Println("Total: ", total)
}
