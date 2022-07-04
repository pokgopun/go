/* https://theweeklychallenge.org/blog/perl-weekly-challenge-148/
TASK #1 › Eban Numbers

Submitted by: [45]Mohammad S Anwar
     __________________________________________________________________

   Write a script to generate all Eban Numbers <= 100.

     An Eban number is a number that has no letter ‘e’ in it when the
     number is spelled in English (American or British).

Example

2, 4, 6, 30, 32 are the first 5 Eban numbers.

*/
package main

import (
	"fmt"
	"regexp"

	"github.com/divan/num2words"
)

func main() {
	eban := []int{}
	re := regexp.MustCompile(`[eE]`)
	for i := 1; i <= 100; i++ {
		if !re.MatchString(num2words.Convert(i)) {
			eban = append(eban, i)
		}
	}
	fmt.Println(eban)
}
