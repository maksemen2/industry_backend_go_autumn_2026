package main

import (
	"errors"
	"fmt"
)

func isMul(a, b int) bool {
	return a%b == 0
}

var ErrZero = errors.New("zero is not allowed")

func fizzBuzz(n int) (string, error) {
	if n == 0 {
		return "", ErrZero
	}

	var result string

	if isMul(n, 3) {
		result += "Fizz"
	}

	if isMul(n, 5) {
		result += "Buzz"
	}

	if result != "" {
		return result, nil
	}

	return fmt.Sprintf("%d", n), nil
}
