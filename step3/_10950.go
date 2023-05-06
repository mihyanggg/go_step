package main

import "fmt"

func main() {
	var test_cnt int
	fmt.Scan(&test_cnt)

	for i := 0; i < test_cnt; i++ {
		var A, B int
		fmt.Scan(&A, &B)
		fmt.Println(A + B)
	}
}
