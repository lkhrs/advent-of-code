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

func compute(operations [][]string) (total int) {
	execute := true
	for _, num := range operations {
		if num[0] == "do()" {
			execute = true
		}
		if num[0] == "don't()" {
			execute = false
		}
		if execute {
			total += stringToI(num[1]) * stringToI(num[2])
		}
	}
	return
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
	regexpMul := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	regexpMulDoDont := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	mul := regexpMul.FindAllStringSubmatch(data, -1)
	mulDoDont := regexpMulDoDont.FindAllStringSubmatch(data, -1)

	// Compute the instructions
	fmt.Println("Total part 1: ", compute(mul))
	fmt.Println("Total part 2: ", compute(mulDoDont))
}
