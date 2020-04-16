package main

import "fmt"

func main() {
	num := 5
	factorial := 1
	for value := 1; value<=num; value++ {
		factorial = factorial * value
	}
	fmt.Printf("Factorial of %d is %d", num, factorial)
}