package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	inp, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < inp; i++ {
		scanner.Scan()
		tokens := strings.Split(scanner.Text(), " ")
		a, _ := strconv.Atoi(tokens[0])
		b, _ := strconv.Atoi(tokens[1])
		writer.WriteString(strconv.Itoa(a+b) + "\n")
	}
}
