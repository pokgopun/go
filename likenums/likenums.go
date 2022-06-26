/*
TASK #2 › Like Numbers
Submitted by: Mohammad S Anwar
You are given positive integers, $m and $n.

Write a script to find total count of integers created using the digits of $m which is also divisible by $n.

Repeating of digits are not allowed. Order/Sequence of digits can’t be altered. You are only allowed to use (n-1) digits at the most. For example, 432 is not acceptable integer created using the digits of 1234. Also for 1234, you can only have integers having no more than three digits.

Example 1:
Input: $m = 1234, $n = 2
Output: 9

Possible integers created using the digits of 1234 are:
1, 2, 3, 4, 12, 13, 14, 23, 24, 34, 123, 124, 134 and 234.

There are 9 integers divisible by 2 such as:
2, 4, 12, 14, 24, 34, 124, 134 and 234.
Example 2:
Input: $m = 768, $n = 4
Output: 3

Possible integers created using the digits of 768 are:
7, 6, 8, 76, 78 and 68.

There are 3 integers divisible by 4 such as:
8, 76 and 68.
*/
package likenums

import (
	"errors"
)

func Count(m, n uint) (c uint, err error) {
	if n == 0 {
		return 0, errors.New("divisor must be greater than zero")
	}
	if m < 10 {
		return 0, errors.New("must have at least two digits")
	}
	// convert m uint 1234 to []byte{1,2,3,4}, m will end up equal 0 and will be reused later to store a composition of its digits
	var (
		bs []byte
		b  byte
	)
	seen := make(map[byte]uint)
	for m > 0 {
		b = byte(m % 10)
		seen[b]++
		if seen[b] > 1 {
			return 0, errors.New("repeating of digits are not allowed")
		}
		bs = append([]byte{b}, bs...)
		m /= 10
	}
	l := len(bs)
	// avoid unnecessary memory allocation for reusable variables in the loop
	var (
		cmb     []byte
		j, cmbl int
	)
	for i := 1; i < l; i++ {
		for _, cmb = range getCombo(i, bs) {
			cmbl = len(cmb)
			// m is reused here
			m, j = 0, 0
			for {
				m += uint(cmb[j])
				j++
				if j < cmbl {
					m *= 10
				} else {
					break
				}
			}
			if m%n == 0 {
				c++
			}
		}
	}
	return c, nil
}
func getCombo(n int, e []byte) (r [][]byte) {
	var c []byte
	var sc []byte
	cTree(n, e, c, func(s []byte) {
		sc = make([]byte, len(s))
		copy(sc, s)
		r = append(r, sc)
	})
	return r
}
func cTree(n int, e []byte, c []byte, f func(s []byte)) {
	if len(c) == n || len(c)+len(e) == n {
		f(append(c, e...)[:n])
	} else {
		for i := 0; len(c)+len(e)-i >= n; i++ {
			cTree(n, e[i+1:], append(c, e[i]), f)
		}
	}
}
