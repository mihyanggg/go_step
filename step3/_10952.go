package main

import "fmt"

func main() {

	for {
		var a, b int
		//fmt.Scanf("%d %d", &a, &b) // 44ms
		fmt.Scan(&a) // 24ms
		fmt.Scan(&b) //

		if a == 0 && b == 0 {
			break
		} else {
			fmt.Println(a + b)
		}
	}

	/* FAILURE
	for {
		scanner := bufio.NewScanner(os.Stdin)
		writer := bufio.NewWriter(os.Stdout)

		scanner.Scan()
		v := scanner.Text()

		var ary [2]int
		index := 0

		for i := range v {
			if string(v[i]) != " " {
				ary[index], _ = strconv.Atoi(string(v[i]))
				index++
				if index == 2 {
					break
				}
			}
		}

		if ary[0] == 0 && ary[1] == 0 {
			break
		} else {
			writer.WriteString(strconv.Itoa(ary[0]+ary[1]) + "\n")
			writer.Flush()
		}
	}
	*/
}
