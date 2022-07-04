/* https://theweeklychallenge.org/blog/perl-weekly-challenge-169/
Task 2: Achilles Numbers

Submitted by: [47]Mohammad S Anwar
     __________________________________________________________________

   Write a script to generate first 20 Achilles Numbers. Please checkout
   [48]wikipedia for more information.

     An Achilles number is a number that is powerful but imperfect (not a
     perfect power). Named after Achilles, a hero of the Trojan war, who
     was also powerful but imperfect.

     A positive integer n is a powerful number if, for every prime factor
     p of n, p^2 is also a divisor.

     A number is a perfect power if it has any integer roots (square
     root, cube root, etc.).

   For example 36 factors to (2, 2, 3, 3) - every prime factor (2, 3) also
   has its square as a divisor (4, 9). But 36 has an integer square root,
   6, so the number is a perfect power.

   But 72 factors to (2, 2, 2, 3, 3); it similarly has 4 and 9 as
   divisors, but it has no integer roots. This is an Achilles number.

Output

 72, 108,  200,  288,  392,  432,  500,  648,  675,  800,  864, 968, 972, 1125,
1152, 1323, 1352, 1372, 1568, 1800
*/
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pokgopun/go/achilles"
)

func main() {
	var n uint = 20
	if len(os.Args) > 1 {
		_, err := fmt.Sscanf(os.Args[1], "%d", &n)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(achilles.NewAchlls(n))
}
