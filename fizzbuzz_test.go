package main

import "testing"

func TestCanSayOne(t *testing.T) {
	one := fizzbuzz(1)

	if one != "1" {
		t.Errorf("Not the right answer. We expected '1' but got '%v'", one)
	}
}
