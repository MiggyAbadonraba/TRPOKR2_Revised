package main

import "errors"

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Деление на ноль")
	}
	return a / b, nil
}
