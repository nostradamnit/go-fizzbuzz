package main

import (
	"math"
	"strconv"
)

func main() {

}

func fizzbuzz(number int) string {
	var answer string

	if isFizz(number) {
		return "Fizz"
	}

	if number == 5 {
		return "Buzz"
	}

	answer = strconv.Itoa(number)
	return answer
}

func isFizz(number int) bool {
	remainder := math.Remainder(float64(number), 3)
	if remainder == 0 {
		return true
	}
	return false
}
