package main

import (
	"fmt"
)

func multiply(a, b int) int {
	return a * b
}

func main() {
	a := 40
	b := 50
	fmt.Printf("a * b = %v * %v = %v", a, b,multiply(a, b))
	fmt.Println()
}