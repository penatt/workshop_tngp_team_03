package main

import "fmt"

func minus(a int, b int) int {
    return a - b
}

func main() {
	a := 40
    b := 50
    result := minus(a, b)
	fmt.Printf("%d - %d = %d\n", a, b, result)
}