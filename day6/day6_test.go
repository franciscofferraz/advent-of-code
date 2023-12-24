package day6

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{[]string{
			"Time:      7  15   30",
			"Distance:  9  40  200",
		}, 288},
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
		input    []string
		expected int
	}{
		{[]string{
			"Time:      7  15   30",
			"Distance:  9  40  200",
		}, 71503},
	}

	for _, test := range tests {
		result := part2(test.input)
		if result != test.expected {
			t.Errorf("For input %q, expected %d, but got %d", test.input, test.expected, result)
		}
	}
}
