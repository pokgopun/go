/* https://theweeklychallenge.org/blog/perl-weekly-challenge-142/
TASK #1 › Divisor Last Digit

Submitted by: [51]Mohammad S Anwar
     __________________________________________________________________

   You are given positive integers, $m and $n.

   Write a script to find total count of divisors of $m having last digit
   $n.

Example 1:

Input: $m = 24, $n = 2
Output: 2

The divisors of 24 are 1, 2, 3, 4, 6, 8 and 12.
There are only 2 divisors having last digit 2 are 2 and 12.

Example 2:

Input: $m = 30, $n = 5
Output: 2

The divisors of 30 are 1, 2, 3, 5, 6, 10 and 15.
There are only 2 divisors having last digit 5 are 5 and 15.

*/
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pokgopun/go/dldcount"
)

func main() {
	s := make([][2]uint, 1)
	if len(os.Args) > 2 {
		_, err := fmt.Sscanf(strings.Join(os.Args[1:], " "), "%d %d", &s[0][0], &s[0][1])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		s[0] = [2]uint{24, 2}
		s = append(s, [2]uint{30, 5})
	}
	for _, v := range s {
		fmt.Printf("Input: m = %d, n = %d\nOutput: %d\n\n", v[0], v[1], dldcount.Count(v[0], v[1]))
	}
}
