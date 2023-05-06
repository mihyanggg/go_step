package main

import "fmt"

func main() {

	var step int

	fmt.Scan(&step)

	for i := range [9]int{} {
		num := i + 1
		fmt.Println(step, "*", num, "=", step*num)
	}
}
