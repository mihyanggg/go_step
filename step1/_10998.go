package main

import "fmt"

func mul(a int, b int) int {
	ret := a * b
	return ret
}

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Print(mul(a, b))
}
