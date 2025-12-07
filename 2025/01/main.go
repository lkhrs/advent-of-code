package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// takes a rotation ("L1" or "R199") and returns a positive or negative int.
func getRotation(rotation string) (int, error) {
	if rotation == "" {
		return 0, nil
	}

	steps, err := strconv.Atoi(rotation[1:])
	if err != nil {
		return 0, err
	}

	switch rotation[0] {
	case 'L':
		return -steps, nil
	case 'R':
		return steps, nil
	}

	return 0, fmt.Errorf("invalid rotation: %s", rotation)
}

// normalizes a dial value to be within 0 and dialSize inclusive.
func normalizeDial(dial int, dialSize int) int {
	modulus := dialSize + 1
	return (dial%modulus + modulus) % modulus
}

func main() {
	dial := 50
	dialSize := 99

	file, err := os.Open("input")
	if err != nil {
		fmt.Println("error opening file:", err)
	}

	zeroCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		rotation, err := getRotation(line)
		if err != nil {
			fmt.Println(err)
			break
		}

		dial = normalizeDial((dial + rotation), dialSize)
		if dial != 0 {
			continue
		}

		zeroCount++
	}
	fmt.Println("password:", zeroCount)
}
