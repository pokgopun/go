/* https://theweeklychallenge.org/blog/perl-weekly-challenge-147/

TASK #2 › Pentagon Numbers

Submitted by: [41]Mohammad S Anwar
     __________________________________________________________________

   Write a script to find the first pair of Pentagon Numbers whose sum and
   difference are also a Pentagon Number.

     Pentagon numbers can be defined as P(n) = n(3n - 1)/2.

Example

The first 10 Pentagon Numbers are:
1, 5, 12, 22, 35, 51, 70, 92, 117 and 145.

P(4) + P(7) = 22 + 70 = 92 = P(8)
but
P(4) - P(7) = |22 - 70| = 48 is not a Pentagon Number.
*/
package main

import (
	"fmt"
	"math"
)

// Set to 0 to have idxPtg() return only positive index, set to 1 the fuction will also return negative index
const minusIdx = 1

func main() {
	i := -20
	maxIdx := 3_000
	for {
		if i > maxIdx-1 {
			break
		}
		pi := calPtg(i)
		j := i
		for {
			j++
			if j > maxIdx {
				break
			}
			pj := calPtg(j)
			if pj+pi < calPtg(j+1) {
				break
			}
			sumIdx := idxPtg(pj + pi)
			if sumIdx != 0 {
				diffIdx := idxPtg(pj - pi)
				if diffIdx != 0 {
					if minusIdx == 1 && i > 0 && j > 0 && sumIdx > 0 && diffIdx > 0 {
						fmt.Println("#################### Positive Index Only ####################")
					}
					fmt.Printf("P(%v) + P(%v) => %v + %v = %v <= P(%v)\n", j, i, pj, pi, pj+pi, sumIdx)
					fmt.Printf("P(%v) - P(%v) => %v - %v = %v <= P(%v)\n", j, i, pj, pi, pj-pi, diffIdx)
					if minusIdx == 1 && i > 0 && j > 0 && sumIdx > 0 && diffIdx > 0 {
						fmt.Println("#############################################################")
					}
					fmt.Print("\n")
					break
				}
			}
		}
		i++
	}
}

func calPtg(i int) int {
	return (3*i - 1) * i / 2
}

// idxPtg check if n is a Pentagon number and return its index if it is, it returns 0 if it is not
func idxPtg(n int) int {
	// p = i(3i - 1)/j2 => 3*i^2 - i -2*p = 0
	// 1 +/- sqrt(1 + 24*p)/6
	for j := 1; j >= 0-minusIdx; j -= 2 {
		i, f := math.Modf((1 + float64(j)*math.Sqrt(1+24*float64(n))) / 6)
		if f == 0 {
			return int(i)
		}
	}
	return 0
}
