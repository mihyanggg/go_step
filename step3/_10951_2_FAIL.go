package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	ary := make([]int, 2)

	for scanner.Scan() {
		line := scanner.Text()

		if scanner.Err() == io.EOF {
			break
		}

		tokens := strings.Split(line, " ")
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
	}

	writer.WriteString(strconv.Itoa(ary[0] + ary[1]))

}
