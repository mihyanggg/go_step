package main

import "fmt"

func inp() (a int, b int) {
	fmt.Scan(&a)
	fmt.Scan(&b)
	return
}

func cal_print(a int, b int) {
	fmt.Println(sum(a, b))
	fmt.Println(min(a, b))
	fmt.Println(mul(a, b))
	fmt.Println(div(a, b))
	fmt.Println(rem(a, b))
}

func sum(a, b int) int {
	return a + b
}

func min(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

// remainder of division
func rem(a, b int) int {
	return a % b
}

func main() {
	a, b := inp()
	cal_print(a, b)
}
