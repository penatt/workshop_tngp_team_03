package main

import (
	"fmt"
	"math"
)

func minus(a int, b int) int {
	return a - b
}

func Divide(a, b int) (int, string) {
	if b != 0 {
		var result int
		result = a / b
		return result, "ok"
	}
	return 0, "Error: Division by zero"
}

func multiply(a, b int) int {
	return a * b
}

func sum(a int, b int) int {
	return a + b
}

func mod(a, b int) int {
	modResult := a % b
	return modResult
}

func sqrt(a float64) float64 {
	resultSqrt := math.Sqrt(a)
	return resultSqrt
}

func main() {

	a := 40
	b := 50
	c := 9.00

	result := minus(a, b)
	result_divide, err := Divide(a, b)
	if err == "ok" {
		fmt.Printf("%d / %d = %d\n", a, b, result_divide)
	} else {
		fmt.Println(err)
	}
	result_multiply := multiply(a, b)
	fmt.Printf("Output:")
	fmt.Printf("%d - %d = %d\n", a, b, result)
	fmt.Printf("a * b = %v * %v = %v\n", a, b, result_multiply)
	fmt.Printf("a + b = %v + %v = %v\n", a, b, sum(a, b))
	fmt.Printf("%v '%%' %v = %v\n", a, b, mod(a, b))
	fmt.Printf("Sqrt(a) = Sqrt(%v) = %v", c, sqrt(c))
}
