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

	if isBuzz(number) {
		return "Buzz"
	}

	answer = strconv.Itoa(number)
	return answer
}

func isFizz(number int) bool {
	return isDivisableBy(3, float64(number))
}

func isBuzz(number int) bool {
	return isDivisableBy(5, float64(number))
}

func isDivisableBy(divisor float64, number float64) bool {
	remainder := math.Remainder(number, divisor)
	if remainder == 0 {
		return true
	}
	return false
}
