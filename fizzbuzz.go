package main

import (
	"math"
	"strconv"
)

func main() {

}

func fizzbuzz(number int) string {
	var answer string

	remainder := math.Remainder(float64(number), 3)
	if remainder == 0 {
		return "Fizz"
	}
	answer = strconv.Itoa(number)
	return answer
}
