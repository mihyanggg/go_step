package main

import "fmt"

func main() {

	var integer int
	fmt.Scan(&integer)

	n := integer / 4
	r := integer % 4

	if r != 0 {
		return
	} else {
		n_cnt := make([]int, n)
		for range n_cnt {
			fmt.Print("long ")
		}
	}
	fmt.Print("int")
}
