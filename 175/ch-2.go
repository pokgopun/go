/* https://theweeklychallenge.org/blog/perl-weekly-challenge-175/
Task 2: Perfect Totient Numbers
Submitted by: Mohammad S Anwar
Write a script to generate first 20 Perfect Totient Numbers. Please checkout wikipedia page for more informations.


Output

3, 9, 15, 27, 39, 81, 111, 183, 243, 255, 327, 363, 471, 729,
2187, 2199, 3063, 4359, 4375, 5571
*/
package main

import (
	"fmt"
	"math/big"
)

var one = big.NewInt(1)

func main() {
	var t Totient
	cntdwn := 20
	for i := int64(2); i < 6000; i++ {
		t.Set(big.NewInt(i))
		if t.IsPerfect() {
			fmt.Println(i)
			cntdwn--
		}
		if cntdwn == 0 {
			break
		}
	}
}

type Totient struct {
	n, val *big.Int
}

func totient(n *big.Int) (r *big.Int) {
	r = big.NewInt(1) // 1 is relatively prime to any number
	//one := big.NewInt(1)
	for i := big.NewInt(2); i.Cmp(n) == -1; i.Add(i, one) {
		if new(big.Int).GCD(nil, nil, i, n).Cmp(one) == 0 {
			r.Add(r, one)
		}
	}
	return r
}

func (t *Totient) Set(n *big.Int) *big.Int {
	t.n = n
	t.val = totient(n)
	return t.val
}

func (t *Totient) IsPerfect() bool {
	sum := big.NewInt(1)
	tmp := new(big.Int).Set(t.val)
	for tmp.Cmp(one) == 1 {
		//fmt.Println("tmp = ", tmp)
		sum.Add(sum, tmp)
		tmp = totient(tmp)
	}
	//fmt.Println("sum =", sum)
	return sum.Cmp(t.n) == 0
}
