package fizzbuzz

import "testing"

func TestNumber1(t *testing.T) {
	result := FizzBuzz(1)
	if result != "1" {
		t.Errorf("FizzBuzz(1) returned %s, expected %s", result, "1")
	}
}

func TestNumber2(t *testing.T) {
	result := FizzBuzz(2)
	if result != "2" {
		t.Errorf("FizzBuzz(2) returned %s, expected %s", result, "2")
	}
}

func TestNumber3(t *testing.T) {
	result := FizzBuzz(3)
	if result != "Fizz" {
		t.Errorf("FizzBuzz(3) returned %s, expected %s", result, "Fizz")
	}
}
