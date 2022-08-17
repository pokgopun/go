/* https://theweeklychallenge.org/blog/perl-weekly-challenge-136/
TASK #2 › Fibonacci Sequence
Submitted by: Mohammad S Anwar
You are given a positive number $n.

Write a script to find how many different sequences you can create using Fibonacci numbers where the sum of unique numbers in each sequence are the same as the given number.

Fibonacci Numbers: 1,2,3,5,8,13,21,34,55,89, …

Example 1
Input:  $n = 16
Output: 4

Reason: There are 4 possible sequences that can be created using Fibonacci numbers
i.e. (3 + 13), (1 + 2 + 13), (3 + 5 + 8) and (1 + 2 + 5 + 8).
Example 2
Input:  $n = 9
Output: 2

Reason: There are 2 possible sequences that can be created using Fibonacci numbers
i.e. (1 + 3 + 5) and (1 + 8).
Example 3
Input:  $n = 15
Output: 2

Reason: There are 2 possible sequences that can be created using Fibonacci numbers
i.e. (2 + 5 + 8) and (2 + 13).
*/
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var samples []uint64
	for _, v := range os.Args[1:] {
		n, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			break
		}
		samples = append(samples, n)
	}
	if len(samples) == 0 {
		samples = []uint64{16, 9, 15}
	}
	for _, n := range samples {
		s := fib(n)
		//fmt.Println(s)
		l := len(s)
		var res uint64
		for r := 2; r <= l; r++ {
			count(int(n), r, s, []uint64{}, &res)
		}
		fmt.Printf("Input: n = %d\nOutput: %d\n\n", n, res)
	}
}

func fib(n uint64) (s []uint64) {
	s = []uint64{0, 1}
	for s[len(s)-1] < n {
		s = append(s, s[len(s)-1]+s[len(s)-2])
	}
	return s[2 : len(s)-1]
}

func count(n, r int, e, c []uint64, res *uint64) uint64 {
	lc, le := len(c), len(e)
	if lc == r || lc+le == r {
		//fmt.Print(n)
		for _, v := range c {
			//fmt.Print("-", v)
			n -= int(v)
		}
		for _, v := range e[:r-lc] {
			//fmt.Print("-", v)
			n -= int(v)
		}
		//fmt.Println("=", n)
		if n == 0 {
			*res += 1
		}
		return 0
	} else {
		for i := 0; i <= lc+le-r; i++ {
			count(n, r, e[i+1:], append(c, e[i]), res)
		}
	}
	return *res
}
