package main

import (
	"fmt"
	"unicode/utf8"
)

func strSplit(str string) string {
	b := []byte(str)
	idx := 0

	for i := 0; i < 50; i++ {
		_, size := utf8.DecodeRune(b[idx:])
		idx += size
	}
	return str[:idx]
}

func main() {

	var id string
	fmt.Scan(&id)

	fmt.Print(id, "??!")
}
