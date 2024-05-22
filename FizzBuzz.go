package fizzbuzz

import "strconv"

func FizzBuzz(inputNumber int) string {
	outputWord := ""
	if inputNumber%3 == 0 {
		outputWord += "Fizz"
	}
	if inputNumber%5 == 0 {
		outputWord += "Buzz"
	}
	if outputWord != "" {
		return outputWord
	}
	return strconv.Itoa(inputNumber)
}
