package fizzbuzz

import "testing"

func TestWhenInput_1_want_1(t *testing.T) {
	fb := FizzBuzz(1)
	expected := "1"
	if fb != expected {
		t.Errorf("Got %s, expected %s", fb, expected)
	}
}

func TestWhenInput_2_want_2(t *testing.T) {
	fb := FizzBuzz(2)
	expected := "2"
	if fb != expected {
		t.Errorf("Got %s, expected %s", fb, expected)
	}
}

func TestWhenInput_3_want_fizz(t *testing.T) {
	fb := FizzBuzz(3)
	expected := "fizz"

	if fb != expected {
		t.Errorf("Got %s, expected %s", fb, expected)
	}
}

func TestWhenInput_5_want_buzz(t *testing.T) {
	fb := FizzBuzz(5)
	expected := "buzz"

	if fb != expected {
		t.Errorf("Got %s, expected %s", fb, expected)
	}
}
