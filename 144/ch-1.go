/* https://theweeklychallenge.org/blog/perl-weekly-challenge-144/
TASK #1 › Semiprime

Submitted by: [56]Mohammad S Anwar
     __________________________________________________________________

   Write a script to generate all Semiprime number <= 100.

   For more information about Semiprime, please checkout the [57]wikipedia
   page.

     In mathematics, a semiprime is a natural number that is the product
     of exactly two prime numbers. The two primes in the product may
     equal each other, so the semiprimes include the squares of prime
     numbers.

Example

10 is Semiprime as 10 = 2 x 5
15 is Semiprime as 15 = 3 x 5

Usage:
To get Semiprime number <= 10
go run ch-1.go 10
4, 6, 9, 10
*/
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pokgopun/go/semiprime"
)

func main() {
	var n uint = 100
	if len(os.Args) > 1 {
		_, err := fmt.Sscanf(os.Args[1], "%d", &n)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(semiprime.NewSemiprime(n))
}
