package main

import "testing"

func TestCanSayOne(t *testing.T) {
	one := fizzbuzz(1)

	if one != "1" {
		t.Errorf("Not the right answer. We expected '1' but got '%v'", one)
	}
}

func TestCanSayFizzFor3(t *testing.T) {
	fizz := fizzbuzz(3)

	if fizz != "Fizz" {
		t.Errorf("We expected 'Fizz', but got '%v'", fizz)
	}
}

func TestCanSayFizzFor6(t *testing.T) {
	fizz := fizzbuzz(6)

	if fizz != "Fizz" {
		t.Errorf("We expected 'Fizz', but got '%v'", fizz)
	}
}

func TestCanSayBuzzFor5(t *testing.T) {
	buzz := fizzbuzz(5)

	if buzz != "Buzz" {
		t.Errorf("We expected 'Buzz', but got '%v'", buzz)
	}
}

func TestCanSayBuzzFor10(t *testing.T) {
	buzz := fizzbuzz(10)

	if buzz != "Buzz" {
		t.Errorf("We expected 'Buzz', but got '%v'", buzz)
	}
}

func TestCanSayFizzBuzzFor15(t *testing.T) {
	fbuzz := fizzbuzz(15)

	if fbuzz != "FizzBuzz" {
		t.Errorf("We expected 'FizzBuzz', but got '%v'", fbuzz)
	}
}
