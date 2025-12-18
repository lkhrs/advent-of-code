package main

import "testing"

func TestGetRotation(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{input: "L10", expected: -10},
		{input: "R5", expected: 5},
		{input: "L0", expected: 0},
		{input: "R0", expected: 0},
		{input: "L500", expected: -500},
		{input: "R500", expected: 500},
		{input: "", expected: 0},
	}

	for _, test := range tests {
		result, err := getRotation(test.input)
		if err != nil {
			t.Errorf("getRotation(%s) returned error: %v", test.input, err)
			continue
		}
		if result != test.expected {
			t.Errorf("getRotation(%s) = %d; expected %d", test.input, result, test.expected)
		}
	}
}

func TestNormalizeDial(t *testing.T) {
	dialSize := 99
	tests := []struct {
		dial     int
		expected int
	}{
		{dial: 0, expected: 0},
		{dial: -1, expected: dialSize},
		{dial: 50, expected: 50},
		{dial: 100, expected: 0},
		{dial: -100, expected: 0},
		{dial: 198, expected: 98},
		{dial: -198, expected: 2},
		{dial: 700, expected: 0},
	}

	for _, test := range tests {
		result := normalizeDial(test.dial, dialSize)
		if result != test.expected {
			t.Errorf("normalizeDial(%d, %d) = %d; expected %d", test.dial, dialSize, result, test.expected)
		}
	}
}

func TestZeroCount(t *testing.T) {
	dial := 50
	dialSize := 99

	steps := []struct {
		rotation     string
		expectedDial int
		expectedLand bool
		expectedPass int
	}{
		{rotation: "L68", expectedDial: 82, expectedLand: false, expectedPass: 1},
		{rotation: "L30", expectedDial: 52, expectedLand: false, expectedPass: 0},
		{rotation: "R48", expectedDial: 0, expectedLand: true, expectedPass: 0},
		{rotation: "L5", expectedDial: 95, expectedLand: false, expectedPass: 0},
		{rotation: "R60", expectedDial: 55, expectedLand: false, expectedPass: 1},
		{rotation: "L55", expectedDial: 0, expectedLand: true, expectedPass: 0},
		{rotation: "L1", expectedDial: 99, expectedLand: false, expectedPass: 0},
		{rotation: "L99", expectedDial: 0, expectedLand: true, expectedPass: 0},
		{rotation: "R14", expectedDial: 14, expectedLand: false, expectedPass: 0},
		{rotation: "L82", expectedDial: 32, expectedLand: false, expectedPass: 1},
		{rotation: "R1000", expectedDial: 32, expectedLand: false, expectedPass: 10},
		{rotation: "L32", expectedDial: 0, expectedLand: true, expectedPass: 0},
		{rotation: "L50", expectedDial: 50, expectedLand: false, expectedPass: 0},
		{rotation: "R101", expectedDial: 51, expectedLand: false, expectedPass: 1},
	}

	for idx, step := range steps {
		landedOnZero, passedZeroCount, newDial := zeroCount(step.rotation, dial, dialSize)

		if landedOnZero != step.expectedLand {
			t.Errorf("step %d (%s): landedOnZero = %t; want %t", idx+1, step.rotation, landedOnZero, step.expectedLand)
		}
		if passedZeroCount != step.expectedPass {
			t.Errorf("step %d (%s): passedZeroCount = %d; want %d", idx+1, step.rotation, passedZeroCount, step.expectedPass)
		}
		if newDial != step.expectedDial {
			t.Errorf("step %d (%s): new dial = %d; want %d", idx+1, step.rotation, newDial, step.expectedDial)
		}

		dial = newDial
	}
}
