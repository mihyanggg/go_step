package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 1; i <= n; i++ {
		for j := n - i; j > 0; j-- {
			writer.WriteString(" ")
		}
		for q := 0; q < i; q++ {
			writer.WriteString("*")
		}
		writer.WriteString("\n")
	}
}
