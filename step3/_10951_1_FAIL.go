package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func input() []int {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	tokens := strings.Split(scanner.Text(), " ")
	ary := make([]int, 2)
	index := 0

	for i := range tokens {
		if tokens[i] != " " {
			ary[index], _ = strconv.Atoi(tokens[i])
			index++
			if index == 2 {
				break
			}
		}
	}

	return ary
}

func output(ary []int) {
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	writer.WriteString(strconv.Itoa(ary[0]+ary[1]) + "\n")
}

func main() {
	for {
		ary := input()
		output(ary)
	}
}
