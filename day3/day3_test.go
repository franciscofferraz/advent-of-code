package day3

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{[]string{
			"467..114..",
			"...*......",
			"..35..633.",
			"......#...",
			"617*......",
			".....+.58.",
			"..592.....",
			"......755.",
			"...$.*....",
			".664.598..",
		}, 4361},
	}

	for _, test := range tests {
		result := part1(test.input)
		if result != test.expected {
			t.Errorf("For input %q, expected %d, but got %d", test.input, test.expected, result)
		}
	}
}
