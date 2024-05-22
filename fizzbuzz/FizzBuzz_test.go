package fizzbuzz

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{"NumberItself1", 1, "1"},
		{"NumberItself2", 2, "2"},
		{"NumberItself4", 4, "4"},
		{"MultipleOf3_3", 3, "Fizz"},
		{"MultipleOf3_9", 9, "Fizz"},
		{"MultipleOf3_99", 99, "Fizz"},
		{"MultipleOf5_5", 5, "Buzz"},
		{"MultipleOf5_10", 10, "Buzz"},
		{"MultipleOf5_100", 100, "Buzz"},
		{"MultipleOf3And5_15", 15, "FizzBuzz"},
		{"MultipleOf3And5_30", 30, "FizzBuzz"},
		{"MultipleOf3And5_60", 60, "FizzBuzz"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FizzBuzz(tt.input)
			if result != tt.expected {
				t.Errorf("FizzBuzz(%d) returned %s, expected %s", tt.input, result, tt.expected)
			}
		})
	}
}
