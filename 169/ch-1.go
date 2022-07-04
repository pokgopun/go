/* https://theweeklychallenge.org/blog/perl-weekly-challenge-169/
Task 1: Brilliant Numbers

Submitted by: [46]Mohammad S Anwar
     __________________________________________________________________

   Write a script to generate first 20 Brilliant Numbers.

     Brilliant numbers are numbers with two prime factors of the same
     length.

   The number should have exactly two prime factors, i.e. it’s the product
   of two primes of the same length.

   For example,
24287 = 149 x 163
24289 = 107 x 227

Therefore 24287 and 24289 are 2-brilliant numbers.
These two brilliant numbers happen to be consecutive as there are no even brilli
ant numbers greater than 14.

Output

4, 6, 9, 10, 14, 15, 21, 25, 35, 49, 121, 143, 169, 187, 209, 221, 247, 253, 289
, 299
*/
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pokgopun/go/brilliant"
)

func main() {
	var n uint = 20
	if len(os.Args) > 1 {
		_, err := fmt.Sscanf(os.Args[1], "%d", &n)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(brilliant.NewBrllnt(n))
}
