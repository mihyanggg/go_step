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
		for j := 0; j < i; j++ {
			writer.WriteString("*")
		}
		writer.WriteString("\n")
	}
}
