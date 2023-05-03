package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	switch {
	case a > 0 && b > 0:
		fmt.Println("1")
	case a < 0 && b > 0:
		fmt.Println("2")
	case a < 0 && b < 0:
		fmt.Println("3")
	case a > 0 && b < 0:
		fmt.Println("4")
	default:
		fmt.Println("0")
	}
}
