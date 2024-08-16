package main

import "fmt"

func minus(a int, b int) int {
    return a + b
}

func main() {
    result := minus(5, 3)
    fmt.Println("The sum is:", result)
}