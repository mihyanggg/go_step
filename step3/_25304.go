package main

import "fmt"

func main() {

	var tot_price, stuff_typ_cnt, stuff_price, stuff_cnt int
	receipt_sum := 0

	fmt.Scan(&tot_price)     // 영수증에 적힌 총 금액 X
	fmt.Scan(&stuff_typ_cnt) // 영수증에 적힌 구매한 물건의 종류의 수 N
	stuff_typ := make([]int, stuff_typ_cnt)

	for range stuff_typ {
		fmt.Scan(&stuff_price, &stuff_cnt) // 각 물건의 가격 a와 개수 b
		tmp := stuff_price * stuff_cnt
		receipt_sum += tmp
	}

	if tot_price == receipt_sum {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
