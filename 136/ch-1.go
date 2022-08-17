/* https://theweeklychallenge.org/blog/perl-weekly-challenge-136/
TASK #1 › Two Friendly
Submitted by: Mohammad S Anwar
You are given 2 positive numbers, $m and $n.

Write a script to find out if the given two numbers are Two Friendly.

Two positive numbers, m and n are two friendly when gcd(m, n) = 2 ^ p where p > 0. The greatest common divisor (gcd) of a set of numbers is the largest positive number that divides all the numbers in the set without remainder.

Example 1
    Input: $m = 8, $n = 24
    Output: 1

    Reason: gcd(8,24) = 8 => 2 ^ 3
Example 2
    Input: $m = 26, $n = 39
    Output: 0

    Reason: gcd(26,39) = 13
Example 3
    Input: $m = 4, $n = 10
    Output: 1

    Reason: gcd(4,10) = 2 => 2 ^ 1
*/
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var samples [][]uint64
	if len(os.Args) > 2 {
		for i := 1; i <= 2; i++ {
			n, err := strconv.ParseUint(os.Args[i], 10, 64)
			if err != nil {
				break
			}
			if len(samples) == 0 {
				samples = append(samples, []uint64{})
			}
			samples[0] = append(samples[0], n)
		}
	}
	if len(samples) == 0 {
		samples = [][]uint64{
			[]uint64{8, 24},
			[]uint64{26, 39},
			[]uint64{4, 10},
		}
	}
	for _, v := range samples {
		fmt.Printf("Input: m = %d, n = %d\nOutput: %t\n\n", v[0], v[1], gcdEuclidean(v[0], v[1])%2 == 0)
	}
}

func gcdEuclidean(a, b uint64) uint64 {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}
