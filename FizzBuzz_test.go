package fizzbuzz

import (
	"testing"
)

func TestFizzBuzzReturningTheNumberItself(t *testing.T) {
	parameters := []struct {
		input    int
		expected string
	}{
		{1, "1"}, {2, "2"}, {4, "4"},
	}

	for i := range parameters {
		result := FizzBuzz(parameters[i].input)
		if result != parameters[i].expected {
			t.Errorf("FizzBuzz() returned %s, expected %s", result, parameters[i].expected)
		}
	}
}

func TestNumber3(t *testing.T) {
	result := FizzBuzz(3)
	if result != "Fizz" {
		t.Errorf("FizzBuzz(3) returned %s, expected %s", result, "Fizz")
	}
}

func TestNumber5(t *testing.T) {
	result := FizzBuzz(5)
	if result != "Buzz" {
		t.Errorf("FizzBuzz(5) returned %s, expected %s", result, "Buzz")
	}
}
