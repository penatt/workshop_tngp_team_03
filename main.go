package main

import "fmt"

func Divide(a, b int) int {
	var result int
	result = a / b
	return result
}

func main() {
	a := 40
	b := 50
	result := Divide(a, b)
	fmt.Println("result :", result)
}
