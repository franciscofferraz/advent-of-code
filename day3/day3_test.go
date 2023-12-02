package day2

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"", 0},
		{"", 0},
		{"", 0},
		{"", 0},
	}

	for _, test := range tests {
		result := part1(test.input)
		if result != test.expected {
			t.Errorf("For input %q, expected %d, but got %d", test.input, test.expected, result)
		}
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"", 0},
		{"", 0},
		{"", 0},
		{"", 0},
	}

	for _, test := range tests {
		result := part2(test.input)
		if result != test.expected {
			t.Errorf("For input %q, expected %d, but got %d", test.input, test.expected, result)
		}
	}
}
