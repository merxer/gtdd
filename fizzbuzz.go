package fizzbuzz

import "strconv"

// FizzBuzz input integer and return string
func FizzBuzz(x int) string {
	if x%3 == 0 {
		return "fizz"
	}
	if x == 5 {
		return "buzz"
	}
	return strconv.Itoa(x)
}
