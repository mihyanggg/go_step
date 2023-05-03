package main

import (
	"fmt"
)

func main() {
	var h, m, for_cook int

	fmt.Scan(&h, &m)
	fmt.Scan(&for_cook)

	m_tmp := m + for_cook

	switch {
	case m_tmp > 59:
		m = m_tmp % 60
		h += (m_tmp / 60)
		if h > 23 {
			h %= 24
		}
	case m_tmp <= 59:
		m = m_tmp
	}

	fmt.Print(h, m)
}

/*
80+40
123 / 60 = 2
123 % 60 = 3
25 % 24 = 1
*/
