/*
https://theweeklychallenge.org/blog/perl-weekly-challenge-171/
Task 1: Abundant Number
Submitted by: Mohammad S Anwar
Write a script to generate first 20 Abundant Odd Numbers.

According to wikipedia,


A number n for which the sum of divisors σ(n) > 2n, or, equivalently, the sum of proper divisors (or aliquot sum) s(n) > n.


For example, 945 is the first Abundant Odd Number.

Sum of divisors:
1 + 3 + 5 + 7 + 9 + 15 + 21 + 27 + 35 + 45 + 63 + 105 + 135 + 189 + 315 = 975
*/
package abundant

import (
	"math"
)

func OddAbundant(n uint) (s []uint) {
	var i uint = 1
	lim := int(n)
	for len(s) < lim {
		if IsAbundant(i) {
			s = append(s, i)
		}
		i += 2
	}
	return s
}

func IsAbundant(n uint) bool {
	if n <= 1 {
		return false
	}
	var d uint = 2
	var step uint = 1
	if n%2 != 0 {
		step = 2
		d++
	}
	lim := uint(math.Floor(math.Sqrt(float64(n))))
	var sum uint = 1
	//var s []uint
	for d <= lim {
		if n%d == 0 {
			sum += d
			//s = append(s, d)
			if n/d != d {
				sum += n / d
				//s = append(s, n/d)
			}
			if sum > n {
				/*
					sort.SliceStable(s, func(i, j int) bool {
						return s[i] < s[j]
					})
					fmt.Println(n, "<", sum, "which is sum of", s)
				*/
				return true
			}
		}
		d += step
	}
	return false
}
