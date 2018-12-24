package fizzbuzz

import "strconv"

// FizzBuzz input integer and return string
func FizzBuzz(x int) string {
	if x%5 == 0 {
		return "buzz"
	}
	if x%3 == 0 {
		return "fizz"
	}
	return strconv.Itoa(x)
}
