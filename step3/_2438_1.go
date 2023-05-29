package main

import "fmt"

func main() {

	var n int
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}
