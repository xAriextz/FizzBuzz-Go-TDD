package fizzbuzz

import "strconv"

func FizzBuzz(inputNumber int) string {
	if inputNumber%3 == 0 {
		return "Fizz"
	}
	if inputNumber%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(inputNumber)
}
