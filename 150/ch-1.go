/* https://theweeklychallenge.org/blog/perl-weekly-challenge-150/
TASK #1 › Fibonacci Words

Submitted by: [40]Mohammad S Anwar
     __________________________________________________________________

   You are given two strings having same number of digits, $a and $b.

   Write a script to generate Fibonacci Words by concatenation of the
   previous two strings. Finally print 51st digit of the first term having
   at least 51 digits.

Example:

    Input: $a = '1234' $b = '5678'
    Output: 7

    Fibonacci Words:

    '1234'
    '5678'
    '12345678'
    '567812345678'
    '12345678567812345678'
    '56781234567812345678567812345678'
    '1234567856781234567856781234567812345678567812345678'

    The 51st digit in the first term having at least 51 digits '1234567856781234
567856781234567812345678567812345678' is 7.

*/
package main

import (
	"fmt"
	"os"
)

func main() {
	a, b := "1234", "5678"
	if len(os.Args) > 2 {
		a, b = os.Args[1], os.Args[2]
	}
	fmt.Printf("Input: $a = '%v' $b = '%v'\n", a, b)
	for len([]rune(b)) < 51 {
		a, b = b, a+b
	}
	fmt.Println("Output:", string([]rune(b)[50]))
}
