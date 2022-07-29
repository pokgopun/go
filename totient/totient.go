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
	"github.com/jbarham/primegen"
)

// processor allow totient computation to use cache which improve performance
type processor struct {
	m map[uint64]uint64 // cache totient value to avoid unnecessary computation
	//pg *primegen.Primegen // being used by Factor() when Euler's product formula is used to compute Totient()
}

// New() return a pointer of processor
func New() *processor {
	// intialize *processor and allow lenghty manipulation of its fields
	/*
		var p processor
		p.m = make(map[uint]uint)
		p.pg = primegen.New()
		return &p
	*/
	return &processor{m: make(map[uint64]uint64)}
	//return &processor{m: make(map[uint64]uint64), pg: primegen.New()}
}

func (p *processor) Totient(n uint64) (r uint64) {
	// https://mathworld.wolfram.com/TotientFunction.html
	// totient(1) = 1, also handle totient(0) at the same time as it can be 0 or 1
	if n <= 1 {
		return 1
	}
	/**/
	// cache only when n is even as perfect totient check will result in n equal even (i.e. n >= 3, totient(n) is always even)
	if n%2 == 0 {
		r, ok := p.m[n]
		if ok {
			return r
		}
	}
	/**/
	/*
		// calcuation using Euler's product formula, need factorization which adds lots of overhead if not probability based implementation
				r = n
				var seen uint64
				for _, v := range p.Factor(n) {
					if seen == v {
						continue
					}
					seen = v
					r /= v
					r *= v - 1
				}
	*/
	/**/
	// calcuation using GCD
	r = 1 // 1 is relatively prime to any number
	for i := uint64(2); i < n; i++ {
		if GCDEuclidean(i, n) == 1 {
			r++
		}
	}
	/**/
	// only cache when n is even as it can be reused by another IsPerfect perform iteration of totient (i.e. n >= 3, totient(n) is always even)
	/**/
	if n%2 == 0 {
		p.m[n] = r
	}
	/**/
	return r
}

// IsPerfect() take an uint64  and check if it is a perfect totient and return bool
func (p *processor) IsPerfect(n uint64) bool {
	// a perfect totient is always an odd number start at 3
	if n < 2 || n%2 == 0 {
		return false
	}
	// initial value of sum is totient(1)=1 which is where we stop the iteration
	var sum uint64 = 1
	// initial value for totient(n) for further iteration
	var t uint64 = p.Totient(n)
	for t > 1 {
		sum += t
		if sum > n {
			return false
		}
		t = p.Totient(t)
	}
	return sum == n
}
func (p *processor) ListPerfect(start uint64, count uint64) (s []uint64) {
	// save and restore the current state of pg if reusing an existing one that may be being used by other method at the same time (i.e. Factor)
	/*
		prev := p.pg.Peek()
		defer func() {
			p.pg.SkipTo(prev)
		}()
		pg := p.pg
	*/
	/**/
	pg := primegen.New()
	pc := newPerfectCandidate(start, pg)
	// iterate, check and store the candidate that is a perfect totient
	var c uint64
	for count > 0 {
		c = pc.next()
		if p.IsPerfect(c) {
			s = append(s, c)
			count--
		}
	}
	/**/
	return s
}

// perfect totient candidate generator to avoid processing even and prime numbers (i.e. except for 3)
type perfectCandidate struct {
	val uint64
	pg  *primegen.Primegen
}

// newPerfectCandidate() takes an uint64 and primegen and then return a pointer of perfectCandidate that use the number to set initial candidate
func newPerfectCandidate(start uint64, pg *primegen.Primegen) *perfectCandidate {
	// 1st perfect totient candidate is 3 and the rest are non-prime odd
	if start < 3 {
		start = 3
	} else if start%2 == 0 {
		start++
	}
	// an initial prime that a perfect totient will never be is the prime number next to the start value
	pg.SkipTo(start)
	for pg.Peek() <= start {
		pg.Next()
	}
	return &perfectCandidate{val: start, pg: pg}
}

// next() return the current candidate and advance the next one
func (pc *perfectCandidate) next() uint64 {
	// advance the odd if it is the current prime, advance the prime if it has been checked against the odd
	for pc.pg.Peek() <= pc.val {
		if pc.val == pc.pg.Peek() {
			pc.val += 2
		}
		pc.pg.Next()
	}
	// if it the current odd is not the current prime and then advance the odd and return the old one
	pc.val += 2
	return pc.val - 2
}

/*
func (p *processor) Factor(n uint64) (s []uint64) {
	switch n {
	case 0, 1:
		return []uint64{}
	case 2, 3:
		return []uint64{n}
	}
	lim := uint64(math.Sqrt(float64(n)))
	prev := p.pg.Peek()
	defer func() {
		p.pg.SkipTo(prev)
	}()
	p.pg.Reset()
	d := p.pg.Next()
	for {
		if n%d == 0 {
			s = append(s, d)
			n /= d
			if n == 1 {
				break
			}
			continue
		}
		d = p.pg.Next()
		if d > lim {
			break
		}
	}
	if n > 1 {
		s = append(s, n)
	}
	return s
}
*/
func GCDEuclidean(a, b uint64) uint64 {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}

/*
func FastPerfect(n uint) (r []uint) {
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
		for len(s) > 0 && s[0] < mot {
			r = append(r, s[0])
			s = s[1:]
			n--
			if n == 0 {
				goto end
			}
		}
		if len(sf) > 0 && sf[0] < mot {
			r = append(r, sf[0])
			sf = sf[1:]
			n--
			if n == 0 {
				break
			}
		}
		r = append(r, mot)
		n--
		if n == 0 {
			break
		}
	}
end:
	return r
}
*/
