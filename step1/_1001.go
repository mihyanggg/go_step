package main

import "fmt"

func inp() (a, b int) {
	fmt.Scan(&a)
	fmt.Scan(&b)
	return a, b
}

func min(a, b int) int {
	res := a - b
	return res
}

func main() {
	a, b := inp()
	res := min(a, b)
	fmt.Println(res)
}
