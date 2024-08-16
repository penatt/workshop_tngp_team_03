package main

import (
	"fmt"
)

func minus(a int, b int) int {
	return a - b
}

func Divide(a, b int) int {
	var result int
	result = a / b
	return result
}

func multiply(a, b int) int {
	return a * b
}

func sum(a int,b int)int{
	return a+b
}

func mod(a, b int) int {
	modResult := a % b
	return modResult
}

func main() {
	a := 40
	b := 50


	

	result := minus(a, b)
	result_divide := Divide(a, b)
	result_multiply := multiply(a, b)
	fmt.Printf("Output:")
	fmt.Printf("%d - %d = %d\n", a, b, result)
	fmt.Printf("%d / %d = %d\n", a, b, result_divide)
	fmt.Printf("a * b = %v * %v = %v\n", a, b,result_multiply)
	fmt.Printf("a + b = %v + %v = %v\n", a,b, sum(a,b))
	fmt.Printf("%v '%%' %v = %v\n", a, b, mod(a,b))
}
