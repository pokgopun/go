/* https://theweeklychallenge.org/blog/perl-weekly-challenge-144/

TASK #2 › Ulam Sequence

Submitted by: [58]Mohammad S Anwar
     __________________________________________________________________

   You are given two positive numbers, $u and $v.

   Write a script to generate Ulam Sequence having at least 10 Ulam
   numbers where $u and $v are the first 2 Ulam numbers.

   For more information about Ulam Sequence, please checkout the
   [59]website.

     The standard Ulam sequence (the (1, 2)-Ulam sequence) starts with U1
     = 1 and U2 = 2. Then for n > 2, Un is defined to be the smallest
     integer that is the sum of two distinct earlier terms in exactly one
     way and larger than all earlier terms.

Example 1

Input: $u = 1, $v = 2
Output: 1, 2, 3, 4, 6, 8, 11, 13, 16, 18

Example 2

Input: $u = 2, $v = 3
Output: 2, 3, 5, 7, 8, 9, 13, 14, 18, 19

Example 3

Input: $u = 2, $v = 5
Output: 2, 5, 7, 9, 11, 12, 13, 15, 19, 23

Usage:
To get a 10 Ulam numbers where u=2 and v=5
go run ch-2.go 2 5 10
2, 5, 7, 9, 11, 12, 13, 15, 19, 23
*/
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pokgopun/go/ulam"
)

func main() {
	var u, v, n uint = 1, 2, 10
	if len(os.Args) > 3 {
		_, err := fmt.Sscanf(strings.Join(os.Args[1:], " "), "%d %d %d", &u, &v, &n)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(ulam.NewUlam(u, v, n))
}
