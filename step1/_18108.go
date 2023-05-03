package main

import "fmt"

func main() {
	var inp_year int // -543
	fmt.Scan(&inp_year)
	if inp_year > 3000 || inp_year < 1000 {
		fmt.Println(" 1000 <= y <= 3000 ")
		return
	}
	fmt.Println(inp_year - 543)
	return
}
