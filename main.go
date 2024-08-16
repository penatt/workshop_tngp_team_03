package main

import "fmt"

func main() {
	mod(40, 50)
}

func mod(a, b int) int {
	modResult := a % b
	fmt.Printf("%v '%%' %v = %v\n", a, b, modResult)
	return modResult
}
