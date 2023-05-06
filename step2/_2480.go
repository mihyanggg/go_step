package main

import "fmt"

func main() {
	var dice [3]int
	var winnings int
	for i := 0; i < 3; i++ {
		fmt.Scan(&dice[i])
		if dice[i] > 6 || dice[i] < 1 {
			fmt.Println("1 <= dice <= 6")
			return
		}
	}

	switch {
	case dice[0] == dice[1] && dice[1] == dice[2]:
		// 같은 눈이 3개가 나오면 10,000원+(같은 눈)×1,000원의 상금
		winnings = 10000 + (dice[0] * 1000)

	case dice[0] == dice[1] || dice[1] == dice[2] || dice[2] == dice[0]:
		// 같은 눈이 2개만 나오는 경우에는 1,000원+(같은 눈)×100원의 상금
		switch {
		case dice[0] == dice[1] || dice[0] == dice[2]:
			winnings = 1000 + (dice[0] * 100)
		case dice[1] == dice[2]:
			winnings = 1000 + (dice[1] * 100)
		}

	default:
		// 모두 다른 눈이 나오는 경우에는 (그 중 가장 큰 눈)×100원의 상금
		var biggest int

		if dice[0] > dice[1] { //0 (1 2)
			if dice[0] > dice[2] { //0 (1 2)
				if dice[1] > dice[2] { //0 1 2
					biggest = dice[0]
				} else { //0 2 1
					biggest = dice[0]
				}
			} else { //2 0 1
				biggest = dice[2]
			}
		} else if dice[1] > dice[2] { //1 (0 2)
			if dice[0] > dice[2] { //1 0 2
				biggest = dice[1]
			} else { //1 2 0
				biggest = dice[1]
			}
		} else { //2 1 0
			biggest = dice[2]
		}

		winnings = biggest * 100
	}

	fmt.Println(winnings)
}
