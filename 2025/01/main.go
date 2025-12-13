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

// returns whether the dial landed on zero, how many times it passed zero, and the new dial position.
func zeroCount(line string, dial, dialSize int) (landedOnZero bool, passedZeroCount, newDial int) {
	rotation, _ := getRotation(line)
	modulus := dialSize + 1
	effectiveRotation := rotation % modulus
	rawDial := dial + effectiveRotation
	newDial = normalizeDial(rawDial, dialSize)
	landedOnZero = newDial == 0

	switch {
	case effectiveRotation > 0 && rawDial >= modulus && newDial != 0:
		passedZeroCount = 1
	case effectiveRotation < 0 && rawDial < 0 && dial != 0:
		passedZeroCount = 1
	}

	return
}

func main() {
	dial := 50
	dialSize := 99
	landedZeroes := 0
	passedZeroes := 0

	file, err := os.Open("input")
	if err != nil {
		fmt.Println("error opening file:", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		landedOnZero, passedZeroCount, newDial := zeroCount(line, dial, dialSize)
		if landedOnZero {
			landedZeroes++
		}
		passedZeroes += passedZeroCount
		dial = newDial
	}

	fmt.Println("part 1 password:", landedZeroes)
	fmt.Println("part 2 password:", landedZeroes+passedZeroes)
}
