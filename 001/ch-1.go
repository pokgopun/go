/* https://theweeklychallenge.org/blog/perl-weekly-challenge-001/
Challenge #1

     Write a script to replace the character ‘e’ with ‘E’ in the string
     ‘Perl Weekly Challenge’. Also print the number of times the
     character ‘e’ is found in the string.

*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Perl Weekly Challenge"
	l := "e"
	fmt.Printf("%q has %q replaced with %q for %v times and resulted in %q\n",
		s, l, strings.ToUpper(l), strings.Count(s, l), strings.ReplaceAll(s, l, strings.ToUpper(l)))
}
