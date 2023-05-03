package main

import "fmt"

func main() {

	var inp [3]int

	for i := range inp {
		fmt.Scan(&inp[i])
		if inp[i] < 2 || inp[i] > 10000 {
			return
		}
	}

	fmt.Println((inp[0] + inp[1]) % inp[2])
	fmt.Println(((inp[0] % inp[2]) + (inp[1] % inp[2])) % inp[2])
	fmt.Println((inp[0] * inp[1]) % inp[2])
	fmt.Println((inp[0] % inp[2]) * (inp[1] % inp[2]) % inp[2])

	return
}
