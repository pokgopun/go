/* https://theweeklychallenge.org/blog/perl-weekly-challenge-143/

TASK #2 › Stealthy Number

Submitted by: [47]Mohammad S Anwar
     __________________________________________________________________

   You are given a positive number, $n.

   Write a script to find out if the given number is Stealthy Number.

     A positive integer N is stealthy, if there exist positive integers
     a, b, c, d such that a * b = c * d = N and a + b = c + d + 1.

Example 1

Input: $n = 36
Output: 1

Since 36 = 4 (a) * 9 (b) = 6 (c) * 6 (d) and 4 (a) + 9 (b) = 6 (c) + 6 (d) + 1.

Example 2

Input: $n = 12
Output: 1

Since 2 * 6 = 3 * 4 and 2 + 6 = 3 + 4 + 1

Example 3

Input: $n = 6
Output: 0

Since 2 * 3 = 1 * 6 but 2 + 3 != 1 + 6 + 1
 You are given a positive number, $n.

*/
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/pokgopun/go/stealthy"
)

func main() {
	var nums []uint
	if len(os.Args) > 1 {
		for _, v := range os.Args[1:] {
			n, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			nums = append(nums, uint(n))
		}
	} else {
		nums = []uint{36, 12, 6}
	}
	for _, v := range nums {
		fmt.Printf("Input: n = %d\nOutput: %t\n\n", v, stealthy.IsStealthy(v))
	}
}
