package main

import "fmt"

func hello() {
	fmt.Println("Hello World!")
}

func sum(num ...int) int {

	tot := 0

	for _, value := range num {
		tot += value
	}

	return tot
}

func main() {
	hello()
}
