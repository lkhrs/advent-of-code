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
	tests := []struct {
		dial     int
		dialSize int
		expected int
	}{
		{dial: 0, dialSize: 99, expected: 0},
		{dial: -1, dialSize: 99, expected: 99},
		{dial: 50, dialSize: 99, expected: 50},
		{dial: 100, dialSize: 99, expected: 0},
		{dial: -100, dialSize: 99, expected: 0},
		{dial: 198, dialSize: 99, expected: 98},
		{dial: -198, dialSize: 99, expected: 2},
		{dial: 700, dialSize: 99, expected: 0},
	}

	for _, test := range tests {
		result := normalizeDial(test.dial, test.dialSize)
		if result != test.expected {
			t.Errorf("normalizeDial(%d, %d) = %d; expected %d", test.dial, test.dialSize, result, test.expected)
		}
	}
}