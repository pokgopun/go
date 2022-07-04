/* https://theweeklychallenge.org/blog/perl-weekly-challenge-168/

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
*/
package main

import (
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/pokgopun/go/homeprime"
)

func main() {
	var s []*big.Int
	if len(os.Args) > 1 {
		for _, v := range os.Args[1:] {
			n, ok := new(big.Int).SetString(v, 10)
			if !ok {
				log.Fatal("failed to set string for big.int")
			}
			s = append(s, n)
		}
	} else {
		for i := int64(2); i <= 64; i++ {
			s = append(s, big.NewInt(i))
		}
	}
	for _, v := range s {
		fmt.Printf("HP(%d) = %v\n", v, homeprime.NewHomePrime(v))
	}
}
