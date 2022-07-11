/* https://theweeklychallenge.org/blog/perl-weekly-challenge-173/
pTask 1: Esthetic Number
Submitted by: Mohammad S Anwar
You are given a positive integer, $n.

Write a script to find out if the given number is Esthetic Number.


An esthetic number is a positive integer where every adjacent digit differs from its neighbour by 1.


For example,

5456 is an esthetic number as |5 - 4| = |4 - 5| = |5 - 6| = 1
120 is not an esthetic numner as |1 - 2| != |2 - 0| != 1
*/
package esthetic

import (
	"math/big"
	"strings"
)

func IsEsthetic(d string) bool {
	n, ok := new(big.Int).SetString(d, 10)
	if !ok || n.Cmp(big.NewInt(10)) == -1 {
		return false
	}
	d = strings.TrimLeft(d, "0")
	for i := 1; i < len(d); i++ {
		if d[i] != d[i-1]-1 && d[i] != d[i-1]+1 {
			return false
		}
	}
	return true
}
