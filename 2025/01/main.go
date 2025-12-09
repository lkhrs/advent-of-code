package main

import (
	"bufio"
	"fmt"
	"math"
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

func zeroCount(line string, dial, dialSize int) (landedOnZero bool, passedZeroCount, newDial int) {
	rotation, _ := getRotation(line)
	rawDial := dial + rotation
	newDial = normalizeDial(rawDial, dialSize)
	landedOnZero = newDial == 0
	passedZeroCount = int(math.Round(float64(rawDial) / float64(dialSize+1)))
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
