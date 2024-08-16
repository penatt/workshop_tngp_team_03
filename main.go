package main

import (
	"fmt"
)

func sum(a int,b int)int{
return a+b
}


func main() {
	a := 40
	b := 50
	fmt.Printf("Output:")
	fmt.Printf("a + b = %v + %v = %v", a,b, sum(a,b))
	fmt.Println()
}