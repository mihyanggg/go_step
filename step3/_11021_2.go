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
		writer.WriteString("Case #" + strconv.Itoa(i+1) + ": " + strconv.Itoa(a+b) + "\n")
	}
}

/*
func main() {
	var inp int
	fmt.Scan(&inp)

	i := 0
	for i < inp {
		var a, b int
		fmt.Scan(&a, &b)
		fmt.Printf("Case #%d: %d\n", i+1, a+b)
		i++
	}
}*/
