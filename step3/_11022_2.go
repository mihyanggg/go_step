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
	test, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < test; i++ {
		scanner.Scan()

		tokens := strings.Split(scanner.Text(), " ")
		a, _ := strconv.Atoi(tokens[0])
		b, _ := strconv.Atoi(tokens[1])

		writer.WriteString("Case #" + strconv.Itoa(i+1) + ": " + strconv.Itoa(a) + " + " + strconv.Itoa(b) + " = " + strconv.Itoa(a+b) + "\n")
		writer.Flush()
	}
}
