package main

import "fmt"

func main() {
	num := 2
	if num == 1 {
		fmt.Printf("1 is a real number")
		return
	}
	if num == 2 {
		fmt.Printf("2 is a prime number")
		return
	}
	for value := 2; value<num; value++ {
		if num%value == 0 {
			fmt.Printf("%d is not a prime number", num)
			return
		}
	}
	fmt.Printf("%d is prime number", num)
}