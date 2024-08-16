package main

import "fmt"

func main() {
	a := 40
	b := 50

	fmt.Printf("%v '%%' %v = %v\n", a, b, mod(a,b))
}

func mod(a, b int) int {
	modResult := a % b
	return modResult
}
