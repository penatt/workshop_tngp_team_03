package main

import "fmt"

func minus(a int, b int) int {
	return a - b
}

func Divide(a, b int) int {
	var result int
	result = a / b
	return result
}

func main() {
	a := 40
	b := 50
	result := minus(a, b)
	result_divide := Divide(a, b)
	fmt.Printf("%d - %d = %d\n", a, b, result)
	fmt.Printf("%d - %d = %d\n", a, b, result_divide)
}
