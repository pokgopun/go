/* https://theweeklychallenge.org/blog/perl-weekly-challenge-163/

Task 2: Summations

Submitted by: [41]Mohammad S Anwar
     __________________________________________________________________

   You are given a list of positive numbers, @n.

   Write a script to find out the summations as described below.

Example 1

Input: @n = (1, 2, 3, 4, 5)
Output: 42

    1 2 3  4  5
      2 5  9 14
        5 14 28
          14 42
             42

The nth Row starts with the second element of the (n-1)th row.
The following element is sum of all elements except first element of previous ro
w.
You stop once you have just one element in the row.

Example 2

Input: @n = (1, 3, 5, 7, 9)
Output: 70

    1 3  5  7  9
      3  8 15 24
         8 23 47
           23 70
              70
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
	if len(argInts.val) > 0 && len(argInts.val) == len(argInts.valid) {
		samples = append(samples, argInts.val)
	} else {
		samples = append(samples, []int{1, 2, 3, 4, 5}, []int{1, 3, 5, 7, 9})
	}
	for _, v := range samples {
		fmt.Printf("Input: n = %v\n", v)
		fmt.Printf("Output: %v\n\n", msaSum(v))
	}
}

func msaSum(s []int) int {
	for len(s) > 1 {
		//fmt.Println(s)
		s = s[1:]
		for i := 1; i < len(s); i++ {
			s[i] += s[i-1]
		}
	}
	return s[0]
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
