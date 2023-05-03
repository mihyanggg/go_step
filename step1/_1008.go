package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func div(a, b float64) float64 {
	return a / b
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	inp, _ := reader.ReadString('\n')
	inp = strings.TrimSpace(inp)
	inps := strings.Split(inp, " ")

	a, _ := strconv.ParseFloat(inps[0], 64)
	b, _ := strconv.ParseFloat(inps[1], 64)

	if 0 < div(a, b) {
		//fmt.Print(div(a, b) * math.Inf(1))
		fmt.Print(div(a, b))
	} else if 0 > div(a, b) {
		//fmt.Print(div(a, b) * math.Inf(-1))
		fmt.Print(div(a, b))
	} else {
		fmt.Print(div(a, b))
	}
}
