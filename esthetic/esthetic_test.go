/* https://theweeklychallenge.org/blog/perl-weekly-challenge-173/
pTask 1: Esthetic Number
Submitted by: Mohammad S Anwar
You are given a positive integer, $n.

Write a script to find out if the given number is Esthetic Number.


An esthetic number is a positive integer where every adjacent digit differs from its neighbour by 1.


For example,

5456 is an esthetic number as |5 - 4| = |4 - 5| = |5 - 6| = 1
120 is not an esthetic numner as |1 - 2| != |2 - 0| != 1
*/
package esthetic

import "testing"

func Test_IsEsthetic(t *testing.T) {
	for d, ans := range map[string]bool{"0": false, "9": false, "10": true, "100": false, "12": true, "120": false, "5456": true, "87898765676565654343": true, "878987656765656543431": false} {
		if res := IsEsthetic(d); res != ans {
			t.Errorf("incorrect result for IsEsthetic(%s): expected %t, got %t", d, ans, res)
		}
	}
}
