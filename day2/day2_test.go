package day2

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"Game 1: 5 red, 1 green, 2 blue; 2 green, 8 blue, 6 red; 8 red, 3 blue, 2 green; 6 red, 1 green, 19 blue; 1 red, 17 blue", 0},
		{"Game 7: 3 blue, 1 red; 3 blue, 10 green; 4 green, 5 blue", 7},
		{"Game 22: 3 red, 1 blue; 3 green, 1 red, 1 blue; 7 green, 2 blue", 22},
		{"Game 27: 4 blue, 15 green; 6 green, 2 blue, 1 red; 9 blue, 10 green, 4 red; 3 red, 3 green, 6 blue; 11 blue, 7 red, 11 green; 6 red, 5 green, 13 blue", 0},
		{"Game 28: 10 blue, 8 red, 10 green; 4 blue, 11 red, 6 green; 8 red, 9 green, 10 blue; 4 red, 9 green, 2 blue", 28},
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
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", 48},
		{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", 12},
		{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", 1560},
		{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", 630},
		{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", 36},
	}

	for _, test := range tests {
		result := part2(test.input)
		if result != test.expected {
			t.Errorf("For input %q, expected %d, but got %d", test.input, test.expected, result)
		}
	}
}
