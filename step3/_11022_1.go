package main

import "fmt"

func main() {
	var test, a, b int

	fmt.Scanln(&test)

	for i := 0; i < test; i++ {
		fmt.Scan(&a)
		fmt.Scan(&b)
		fmt.Printf("Case #%d: %d + %d = %d\n", i+1, a, b, a+b)
	}
}
