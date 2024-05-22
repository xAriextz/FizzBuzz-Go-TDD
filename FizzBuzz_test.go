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

func TestMultiplesOf3(t *testing.T) {
	parameters := []struct {
		input    int
		expected string
	}{
		{3, "Fizz"}, {9, "Fizz"}, {99, "Fizz"},
	}

	for i := range parameters {
		result := FizzBuzz(parameters[i].input)
		if result != parameters[i].expected {
			t.Errorf("FizzBuzz() returned %s, expected %s", result, parameters[i].expected)
		}
	}
}

func TestMultiplesOf5(t *testing.T) {
	parameters := []struct {
		input    int
		expected string
	}{
		{5, "Buzz"}, {10, "Buzz"}, {100, "Buzz"},
	}

	for i := range parameters {
		result := FizzBuzz(parameters[i].input)
		if result != parameters[i].expected {
			t.Errorf("FizzBuzz() returned %s, expected %s", result, parameters[i].expected)
		}
	}
}

func TestMultiplesOf3And5(t *testing.T) {
	parameters := []struct {
		input    int
		expected string
	}{
		{15, "FizzBuzz"}, {30, "FizzBuzz"}, {60, "FizzBuzz"},
	}

	for i := range parameters {
		result := FizzBuzz(parameters[i].input)
		if result != parameters[i].expected {
			t.Errorf("FizzBuzz() returned %s, expected %s", result, parameters[i].expected)
		}
	}
}
