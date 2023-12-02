package day1

import "testing"

func TestCalibrate(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"two1nine", 29},
		{"onerbfkf4threeone", 11},
		{"seveneightnine", 79},
		{"invalid number", 0},
		{"two23abc45oneight", 28},
	}

	for _, test := range tests {
		result := calibrate(test.input)
		if result != test.expected {
			t.Errorf("For input %q, expected %d, but got %d", test.input, test.expected, result)
		}
	}
}
