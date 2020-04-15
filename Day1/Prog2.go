package main

import "fmt"

func main() {
	lines := 4
	for line := 1; line<=lines; line++ {
		for  item := 1; item<=line; item++ {
			fmt.Printf("*\t")
		}
		fmt.Printf("\n")
	}
}