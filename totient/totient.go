/* https://theweeklychallenge.org/blog/perl-weekly-challenge-175/
Task 2: Perfect Totient Numbers
Submitted by: Mohammad S Anwar
Write a script to generate first 20 Perfect Totient Numbers. Please checkout wikipedia page for more informations.


Output

3, 9, 15, 27, 39, 81, 111, 183, 243, 255, 327, 363, 471, 729,
2187, 2199, 3063, 4359, 4375, 5571
*/
package totient

import (
	"math/big"
)

type processor struct {
	m             map[uint]uint
	one, min, max *big.Int
	mot           uint
}

func New() processor {
	return processor{m: make(map[uint]uint), one: big.NewInt(1), mot: 1}
}

func (p *processor) Totient(n uint) (r uint) {
	if n == 1 {
		return 1
	}
	r, ok := p.m[n]
	if ok {
		return r
	}
	r = 1 // 1 is relatively prime to any number
	//m := big.NewInt(int64(n))
	p.max = big.NewInt(int64(n))
	//for i := big.NewInt(2); i.Cmp(m) == -1; i.Add(i, p.one) {
	for p.min = big.NewInt(2); p.min.Cmp(p.max) == -1; p.min.Add(p.min, p.one) {
		//if new(big.Int).GCD(nil, nil, i, m).Cmp(p.one) == 0 {
		if new(big.Int).GCD(nil, nil, p.min, p.max).Cmp(p.one) == 0 {
			r++
		}
	}
	p.m[n] = r
	return r
}

func (p *processor) IsPerfect(n uint) bool {
	if n < 2 || n%2 == 0 {
		return false
	}
	// check if n is multiple of 3 as well as isPrime(4*3^k + 1)
	p.mot = 1
	if n%3 == 0 {
		/**/
		pp := n / 3
		if big.NewInt(int64(pp)).ProbablyPrime(0) {
			pp--
			if pp%4 == 0 {
				pp /= 4
			} else {
				pp = 0
			}
		} else {
			pp = 0
		}
		/**/
		for p.mot < n {
			/**/
			if pp != 0 && pp*p.mot == n {
				return true
			}
			/**/
			p.mot *= 3
		}
		if p.mot == n {
			return true
		}
	}
	var sum, tmp uint = 1, p.Totient(n)
	for tmp > 1 {
		sum += tmp
		tmp = p.Totient(tmp)
	}
	return sum == n
}
