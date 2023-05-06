package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var inp int
	fmt.Scan(&inp)

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	i := 0
	for i < inp {

		sum := 0
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		nums := strings.Split(input, " ")

		if 0 < len(nums) {
			for i := range nums {
				num, _ := strconv.Atoi(nums[i])
				sum += num
			}
		}
		fmt.Fprintf(writer, "%d\n", sum)
		writer.Flush()

		i++
	}

	writer.Flush()

}
