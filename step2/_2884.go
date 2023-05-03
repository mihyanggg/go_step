package main

import (
	"fmt"
)

func main() {
	var h, m int
	fmt.Scan(&h, &m)

	tmp := m - 45

	switch {
	case tmp == 0 || tmp > 0:
		m = tmp
	case tmp < 0:
		m = 60 + tmp
		h -= 1
		if h < 0 {
			h = 23
		}
	}
	fmt.Print(h, m)
}
