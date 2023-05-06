package main

import "fmt"

func main() {
	var num int
	sum := 0
	fmt.Scan(&num)

	for i := 1; i <= num; i++ {
		sum = i + sum
	}
	fmt.Println(sum)
}
