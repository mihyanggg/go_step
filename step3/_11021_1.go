package main

import "fmt"

func main() {
	var inp int
	fmt.Scan(&inp)

	i := 0
	for i < inp {
		var a, b int
		fmt.Scan(&a, &b)
		fmt.Printf("Case #%d: %d\n", i+1, a+b)
		i++
	}
}
