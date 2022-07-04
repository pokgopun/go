/* https://theweeklychallenge.org/blog/perl-weekly-challenge-168/
Task 1: Perrin Prime

Submitted by: [45]Roger Bell_West
     __________________________________________________________________

   The Perrin sequence is defined to start with [3, 0, 2]; after that,
   term N is the sum of terms N-2 and N-3. (So it continues 3, 2, 5, 5, 7,
   ….)

     A Perrin prime is a number in the Perrin sequence which is also a
     prime number.

   Calculate the first 13 Perrin Primes.

   f(13) = [2, 3, 5, 7, 17, 29, 277, 367, 853, 14197, 43721, 1442968193,
   792606555396977]

Task 2: Home Prime

Submitted by: [46]Mohammad S Anwar
     __________________________________________________________________

   You are given an integer greater than 1.

   Write a script to find the home prime of the given number.

   In number theory, the home prime HP(n) of an integer n greater than 1
   is the prime number obtained by repeatedly factoring the increasing
   concatenation of prime factors including repetitions.

   Further information can be found on [47]Wikipedia and [48]OEIS.

Example

   As given in the Wikipedia page,
HP(10) = 773, as 10 factors as 2×5 yielding HP10(1) = 25, 25 factors as 5×5 yiel
ding HP10(2) = HP25(1) = 55, 55 = 5×11 implies HP10(3) = HP25(2) = HP55(1) = 511
, and 511 = 7×73 gives HP10(4) = HP25(3) = HP55(2) = HP511(1) = 773, a prime num
ber.
 The Perrin sequence is defined to start with [3, 0, 2]; after that, term N is the sum of terms N-2 and N-3. (So it continues 3, 2, 5, 5, 7, ….)

A Perrin prime is a number in the Perrin sequence which is also a prime number.

Calculate the first 13 Perrin Primes.

f(13) = [2, 3, 5, 7, 17, 29, 277, 367, 853, 14197, 43721, 1442968193, 792606555396977]
*/
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pokgopun/go/perrinprime"
)

func main() {
	n := 13
	if len(os.Args) > 1 {
		_, err := fmt.Sscanf(os.Args[1], "%d", &n)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("f(%d) = [%v]\n", n, perrinprime.NewPerrinPrime(n))
}
