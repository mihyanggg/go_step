package main

import (
	"fmt"
)

func main() {
	var A, B int

	fmt.Scan(&A, &B)

	if A < -10000 || B < -10000 || A > 10000 || B > 10000 {
		return
	}

	if A > B {
		fmt.Println(">")
	} else if A < B {
		fmt.Println("<")
	} else {
		fmt.Println("==")
	}

	return
}
