package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		for j := n - i; j > 0; j-- {
			fmt.Printf(" ")
		}
		for q := 0; q < i; q++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}
