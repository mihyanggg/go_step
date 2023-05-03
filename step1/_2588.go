package main

import "fmt"

func main() {
	var inp [2]int

	for i := range inp {
		var tmp int
		fmt.Scan(&tmp)
		if tmp < 100 || tmp > 999 {
			return
		}
		inp[i] = tmp
	}

	var res [3]int
	tmp := inp[1]

	for i := range res {
		res[i] = tmp % 10
		tmp = tmp / 10
	}

	for i := range res {
		fmt.Println(inp[0] * res[i])
	}
	fmt.Println(inp[0] * inp[1])

}
