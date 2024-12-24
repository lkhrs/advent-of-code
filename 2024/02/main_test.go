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
		input    []string
		expected bool
	}{
		{[]string{"4", "5", "6"}, true},
		{[]string{"4", "8", "15", "16", "23", "42"}, false},
		{[]string{"1", "2", "3", "4", "5"}, true},
		{[]string{"1", "4", "7", "10"}, true},
		{[]string{"10", "9", "8", "7"}, true},
		{[]string{"10", "7", "4", "1"}, true},
		{[]string{"1", "2", "2", "3"}, false},
		{[]string{"1", "2", "5"}, true},
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
		dampener bool
		expected int
	}{
		{
			name: "All safe reports without dampener",
			reports: [][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			},
			dampener: false,
			expected: 3,
		},
		{
			name: "Some unsafe reports without dampener",
			reports: [][]string{
				{"1", "2", "3"},
				{"4", "8", "6"},
				{"7", "8", "9"},
			},
			dampener: false,
			expected: 2,
		},
		{
			name: "All safe reports with dampener",
			reports: [][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			},
			dampener: true,
			expected: 3,
		},
		{
			name: "Some unsafe reports with dampener",
			reports: [][]string{
				{"1", "2", "3"},
				{"4", "8", "6"},
				{"7", "8", "9"},
			},
			dampener: true,
			expected: 3,
		},
		{
			name: "Reports with large differences",
			reports: [][]string{
				{"1", "2", "10"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			},
			dampener: true,
			expected: 3,
		},
		{
			name: "Reports with equal values",
			reports: [][]string{
				{"1", "1", "1"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			},
			dampener: true,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := safeReports(tt.reports, tt.dampener)
			if got != tt.expected {
				t.Errorf("safeReports() = %v, want %v", got, tt.expected)
			}
		})
	}
}
