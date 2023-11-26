package main

import (
	"strconv"
)

func main() {

}

func fizzbuzz(number int) string {
	var answer string

	if number == 3 {
		return "Fizz"
	}
	answer = strconv.Itoa(number)
	return answer
}
