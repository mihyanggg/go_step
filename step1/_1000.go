package main

import "fmt"

func inp() (int, int) {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	return a, b
}

func sum(a int, b int) int {
	tot := a + b
	return tot
}

func main() {
	a, b := inp()
	tot := sum(a, b)
	fmt.Println(tot)
}
