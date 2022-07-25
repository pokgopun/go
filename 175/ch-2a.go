package main

import (
	"fmt"
	"math/big"
)

func main() {
	cntdwn := 20
	var mot int64 = 1
	var s []int64
	var j int64
	for {
		j = 4*mot + 1
		if big.NewInt(j).ProbablyPrime(0) {
			s = append(s, j*3)
		}
		mot *= 3
		if len(s) > 0 && s[0] < mot {
			fmt.Println(s[0])
			s = s[1:]
			cntdwn--
			if cntdwn == 0 {
				break
			}
		}
		fmt.Println(mot)
		cntdwn--
		if cntdwn == 0 {
			break
		}
	}
}
