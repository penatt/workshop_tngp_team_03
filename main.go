package main

import "fmt"

func main() {
	fmt.Println("Output: ",mod(40,50))
}

func mod(a,b int) int {
	modResult := a%b
	return modResult
}