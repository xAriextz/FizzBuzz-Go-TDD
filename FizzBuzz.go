package fizzbuzz

import "strconv"

func FizzBuzz(inputNumber int) string {
	if inputNumber%3 == 0 {
		return "Fizz"
	}
	return strconv.Itoa(inputNumber)
}
