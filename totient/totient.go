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
	fp            map[uint]bool
}

func New(n uint) processor {
	fp := make(map[uint]bool)
	for _, v := range fastperfect(n) {
		fp[v] = true
	}
	return processor{m: make(map[uint]uint), one: big.NewInt(1), mot: 1, fp: fp}
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
	// check for fastperfect
	if p.fp[n] {
		return true
	}
	// check if n is multiple of 3 as well as isPrime(4*3^k + 1)
	/*
		p.mot = 1
		if n%3 == 0 {
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
			for p.mot < n {
				if pp != 0 && pp*p.mot == n {
					return true
				}
				p.mot *= 3
			}
			if p.mot == n {
				return true
			}
		}
	*/
	var sum, tmp uint = 1, p.Totient(n)
	for tmp > 1 {
		sum += tmp
		if sum > n {
			return false
		}
		tmp = p.Totient(tmp)
	}
	return sum == n
}
func fastperfect(n uint) (r []uint) {
	var sf []uint
	var mot uint = 3
	for _, v := range []uint{5, 17, 257, 65537} {
		mot *= v
		sf = append(sf, mot)
	}
	if n > 50 {
		n = 50
	}
	mot = 1
	var s []uint
	var j uint
	for {
		j = 4*mot + 1
		if big.NewInt(int64(j)).ProbablyPrime(0) {
			if len(sf) > 0 && sf[0] == j*3 {
				sf = sf[1:]
			}
			if len(sf) > 0 && sf[0] < j*3 && sf[0] > mot*3 {
				s = append(s, sf[0])
				sf = sf[1:]
			}
			s = append(s, j*3)
		}
		mot *= 3
		//if len(s) > 0 && s[0] < mot {
		for len(s) > 0 && s[0] < mot {
			//fmt.Println(s[0])
			r = append(r, s[0])
			s = s[1:]
			n--
			if n == 0 {
				break
			}
		}
		if len(sf) > 0 && sf[0] < mot {
			//fmt.Println(sf[0])
			r = append(r, sf[0])
			sf = sf[1:]
			n--
			if n == 0 {
				break
			}
		}
		//fmt.Println(mot)
		r = append(r, mot)
		n--
		if n == 0 {
			break
		}
	}
	return r
}
