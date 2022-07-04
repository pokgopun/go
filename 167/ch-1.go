/* https://theweeklychallenge.org/blog/perl-weekly-challenge-167/
Task 1: Circular Prime

Submitted by: [41]Mohammad S Anwar
     __________________________________________________________________

   Write a script to find out first 10 circular primes having at least 3
   digits (base 10).

   Please checkout [42]wikipedia for more information.

     A circular prime is a prime number with the property that the number
     generated at each intermediate step when cyclically permuting its
     (base 10) digits will also be prime.

Output

113, 197, 199, 337, 1193, 3779, 11939, 19937, 193939, 199933

*/
package main

import (
	"io"
	"math"
	"os"
	"strconv"
)

func main() {
	var n uint = 1_000_000
	m := pmap(n)
	var str string
	var j, k int
	for i := 100; i <= int(n); i++ {
		if m[i] {
			j = 0
			for k, _ = range calt(uint(i)) {
				if !m[k] {
					j++
				}
				m[k] = false
			}
			if j == 0 {
				str += strconv.Itoa(i) + ", "
			}
		}
	}
	io.WriteString(os.Stdout, str[:len(str)-2]+"\n")
}

func calt(n uint) map[int]struct{} {
	str := strconv.Itoa(int(n))
	l := len(str)
	c := make(map[int]struct{}, l-1)
	var j int
	for i := 1; i < l; i++ {
		str = str[1:] + string(str[0])
		j, _ = strconv.Atoi(str)
		c[j] = struct{}{}
	}
	return c
}

func pmap(n uint) []bool {
	m := make([]bool, int(n)+2)
	for i := 2; i <= int(n); i++ {
		m[i] = true
	}
	for i := 2; float64(i) <= math.Sqrt(float64(n)); i++ {
		j := i * i
		for j <= int(n) {
			m[j] = false
			j += i
		}

	}
	return m
}
