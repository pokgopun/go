/*
https://theweeklychallenge.org/blog/perl-weekly-challenge-163/
Task 1: Sum Bitwise Operator

Submitted by: [40]Mohammad S Anwar
     __________________________________________________________________

   You are given list positive numbers, @n.

   Write script to calculate the sum of bitwise & operator for all unique
   pairs.

Example 1

Input: @n = (1, 2, 3)
Output: 3

Since (1 & 2) + (2 & 3) + (1 & 3) => 0 + 2 + 1 =>  3.

Example 2

Input: @n = (2, 3, 4)
Output: 2

Since (2 & 3) + (2 & 4) + (3 & 4) => 2 + 0 + 0 =>  2.

*/
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var samples [][]int
	argInts := newArgInts()
	if len(argInts.val) > 1 && len(argInts.val) == len(argInts.valid) {
		samples = append(samples, argInts.val)
	} else {
		samples = append(samples, []int{1, 2, 3}, []int{2, 3, 4})
	}
	for _, v := range samples {
		fmt.Printf("Input: n = %v\nOutput: %v\n\n", v, sumBitwise(v))
	}
}

func sumBitwise(s []int) (sum int) {
	for len(s) > 0 {
		i := s[0]
		s = s[1:]
		for _, j := range s {
			sum += i & j
		}
	}
	return sum
}

type argInts struct {
	valid map[int]bool
	val   []int
}

func newArgInts() (ais argInts) {
	ais.valid = map[int]bool{}
	for i, v := range os.Args[1:] {
		r, err := strconv.Atoi(v)
		if err == nil {
			ais.valid[i] = true
		}
		ais.val = append(ais.val, r)
	}
	return ais
}
