package main

import (
	"fmt"
)

func main() {
	var year int
	fmt.Scan(&year)

	switch {
	case year%4 == 0:
		switch {
		case year%100 != 0:
			fallthrough
		case year%400 == 0:
			fmt.Print("1")
		default:
			fmt.Print("0")
		}
	default:
		fmt.Print("0")
	}

}
