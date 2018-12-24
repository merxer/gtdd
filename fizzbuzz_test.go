package fizzbuzz

import "testing"

func TestWhenInput_1_want_1(t *testing.T) {
	fb := FizzBuzz(1)
	expected := "1"
	if fb != expected {
		t.Errorf("Got %s, expected %s", fb, expected)
	}
}