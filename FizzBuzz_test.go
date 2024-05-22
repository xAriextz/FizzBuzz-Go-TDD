package fizzbuzz

import "testing"

func TestNumber1(t *testing.T) {
	result := FizzBuzz(1)
	if result != 1 {
		t.Errorf("FizzBuzz(1) returned %d, expected %d", result, 1)
	}
}
