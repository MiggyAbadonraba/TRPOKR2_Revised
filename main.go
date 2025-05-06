package main

import "fmt"

func main() {
	a, b := 10.0, 3.0

	fmt.Println("Сложение:", Add(a, b))
	fmt.Println("Вычитание:", Sub(a, b))

	div := Divide(a, 2)

	fmt.Println("Деление:", div)

	mul := Multiply(b, 5)

	fmt.Println("Умножение:", mul)

	pow := Power(a, 3)

	fmt.Println("Степень:", pow)
}
