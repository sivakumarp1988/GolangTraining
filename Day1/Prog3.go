package main

import "fmt"

func main() {
	lines := 5
	value := 1
	for line := 1; line<=lines; line++ {
		for  item := 1; item<=line; item++ {
			fmt.Printf("%d\t",value)
			value=value+1
		}
		fmt.Printf("\n")
	}
}