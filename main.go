package main

import "fmt"

func main() {
	mod(40, 50)
}

func mod(a, b int) int {
	modResult := b % a
	fmt.Printf("%v '%%' %v = %v\n", b, a, modResult)
	return modResult
}
