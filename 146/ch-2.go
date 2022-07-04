/* https://theweeklychallenge.org/blog/perl-weekly-challenge-146/

TASK #2 › Curious Fraction Tree

Submitted by: [45]Mohammad S Anwar
     __________________________________________________________________

   Consider the following Curious Fraction Tree:

   Curious Fraction Tree

   You are given a fraction, member of the tree created similar to the
   above sample.

   Write a script to find out the parent and grandparent of the given
   member.

Example 1:

    Input: $member = '3/5';
    Output: parent = '3/2' and grandparent = '1/2'

Example 2:

    Input: $member = '4/3';
    Output: parent = '1/3' and grandparent = '1/2'
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	var sample []string
	if len(os.Args) > 1 {
		sample = os.Args[1:]
	} else {
		sample = []string{
			"redivider",
			"deific",
			"rotors",
			"challenge",
			"champion",
			"christmas",
		}
	}
	fmt.Println(sample)
}
