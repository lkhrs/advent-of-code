package main

import (
	"reflect"
	"testing"
)

func TestPositive(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{input: -5, expected: 5},
		{input: 5, expected: 5},
		{input: 0, expected: 0},
	}

	for _, test := range tests {
		result := positive(test.input)
		if result != test.expected {
			t.Errorf("positive(%d) = %d; expected %d", test.input, result, test.expected)
		}
	}
}

func TestStrToIntSlice(t *testing.T) {
	tests := []struct {
		input    []string
		expected []int
	}{
		{input: []string{"4", "8", "15", "16", "23", "42"}, expected: []int{4, 8, 15, 16, 23, 42}},
	}

	for _, test := range tests {
		result := sliceStrToInt(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("strToIntSlice(%s) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

func TestAcceptableRange(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{input: []int{7, 6, 4, 2, 1}, expected: true},
		{input: []int{1, 2, 7, 8, 9}, expected: false},
		{input: []int{9, 7, 6, 2, 1}, expected: false},
		{input: []int{1, 3, 2, 4, 5}, expected: false},
		{input: []int{8, 6, 4, 4, 1}, expected: false},
		{input: []int{1, 3, 6, 7, 9}, expected: true},
		{input: []int{15, 16, 18, 20, 23}, expected: true},
	}

	for _, test := range tests {
		result := acceptableRange(test.input)
		if result != test.expected {
			t.Errorf("acceptableRange(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

func TestSafeReports(t *testing.T) {
	tests := []struct {
		name     string
		reports  [][]string
		expected int
	}{
		{
			name:     "All safe reports",
			reports:  [][]string{{"2", "3", "4"}, {"5", "6", "7"}, {"15", "16", "17"}},
			expected: 3,
		},
		{
			name:     "No safe reports",
			reports:  [][]string{{"481"}, {"1623"}, {"2345"}, {"4231"}},
			expected: 0,
		},
		{
			name:     "Mixed safe and unsafe reports",
			reports:  [][]string{{"4", "5"}, {"8", "14"}, {"15", "16"}, {"1623"}, {"2345"}},
			expected: 2,
		},
		{
			name:     "Empty reports",
			reports:  [][]string{},
			expected: 0,
		},
		{
			name:     "Single safe report",
			reports:  [][]string{{"4", "5"}},
			expected: 1,
		},
		{
			name:     "Single unsafe report",
			reports:  [][]string{{"4", "3", "2", "2"}},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := safeReports(tt.reports)
			if result != tt.expected {
				t.Errorf("safeReports(%v) = %d; expected %d", tt.reports, result, tt.expected)
			}
		})
	}
}
