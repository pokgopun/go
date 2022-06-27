/*
You are given an integer greater than 1.

Write a script to find the home prime of the given number.

In number theory, the home prime HP(n) of an integer n greater than 1 is the prime number obtained by repeatedly factoring the increasing concatenation of prime factors including repetitions.
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	/*
		var s []uint
		if len(os.Args) > 1 {
			for _, v := range os.Args[1:] {
				n, err := strconv.ParseUint(v, 10, 64)
				if err != nil {
					log.Fatal(err)
				}
				s = append(s, uint(n))
			}
		} else {
			for i := 2; i <= 47; i++ {
				s = append(s, uint(i))
			}
		}
		for _, v := range s {
			fmt.Printf("HP(%d) = %v\n", v, newHomePrime(v))
		}
	*/
	/**/
	var n, o uint
	if len(os.Args) > 1 {
		c, err := fmt.Sscanf(strings.Join(os.Args[1:], " "), "%d %d", &n, &o)
		if c < 1 {
			log.Fatal(err)
		}
	}
	ch, f := newPrimeGenerator(0).listPrime(int(n))
	defer f()
	w := bufio.NewWriter(os.Stdout)
	for v := range ch {
		w.WriteString(strconv.Itoa(v) + "\n")
	}
	w.Flush()
	/**/
}

type factorizer interface {
	factorize(n *big.Int) []*big.Int
}

type homePrime struct {
	big.Int
	factorizer
}

type factorizerFunc func(n *big.Int) []*big.Int

func (f factorizerFunc) factorize(n *big.Int) []*big.Int {
	return f(n)
}

func newHomePrime(n uint) (hp homePrime) {
	if n < 2 {
		log.Fatal("input must be greater than 1")
	}
	hp.SetInt64(int64(n))
	//hp.SetString(strconv.FormatUint(uint64(n), 10), 10)
	hp.factorizer = factorizerFunc(factor)
	for !hp.ProbablyPrime(0) {
		var str string
		for _, v := range hp.factorize(&hp.Int) {
			str += v.String()
		}
		hp.SetString(str, 10)
	}
	return hp
}

func (hp homePrime) String() string {
	return hp.Int.String()
}
func factor(n *big.Int) (s []*big.Int) {
	if n.Cmp(big.NewInt(1)) == 1 {
		if n.ProbablyPrime(0) {
			return []*big.Int{n}
		}
		ch, f := newPrimeGenerator(5_000_000).listPrime(int(big.NewInt(0).Sqrt(n).Int64()))
		defer f()
		/*
			for v := range ch {
				d := big.NewInt(int64(v))
				if big.NewInt(0).Mul(d, d).Cmp(n) == 1 {
					break
				}
				for big.NewInt(0).Mod(n, d).Cmp(big.NewInt(0)) == 0 {
					s = append(s, d)
					n.Div(n, d)
				}
			}
			if n.Cmp(big.NewInt(1)) == 1 {
				s = append(s, n)
			}
		*/
		d := big.NewInt(int64(<-ch))
		for {
			if big.NewInt(0).Mod(n, d).Cmp(big.NewInt(0)) != 0 {
				d.SetInt64(int64(<-ch))
			} else {
				s = append(s, big.NewInt(0).Set(d))
				//fmt.Println(s)
				n.Div(n, d)
				//fmt.Println(n)
				if n.Cmp(big.NewInt(1)) == 0 {
					break
				} else if n.ProbablyPrime(0) {
					s = append(s, n)
					break
				}
			}
		}
	}
	return s
}

type primeGenerator struct {
	bLimit int
	m      []bool
}

func newPrimeGenerator(bLimit uint) *primeGenerator {
	var pg primeGenerator
	if bLimit == 0 {
		bLimit = 5_000_000
	}
	pg.bLimit = int(bLimit)
	return &pg
}

func (pg *primeGenerator) listPrime(n int) (<-chan int, func()) {
	o := 0
	if n > pg.bLimit {
		o = int(math.Floor(math.Sqrt(float64(n))))
		ch, f := pg.genPrime(o, 0)
		defer f()
		oo := o/2 + o%2
		pg.m = make([]bool, oo)
		for i := 0; i < oo; i++ {
			pg.m[i] = true
		}
		n -= o
		for v := range ch {
			pg.m[(v-1)/2] = false
		}
	}
	q := n / pg.bLimit
	r := n % pg.bLimit
	out := make(chan int)
	done := make(chan struct{})
	go func() {
		var ch <-chan int
		var f func()
		for k, v := range pg.m {
			if v {
				continue
			}
			select {
			case <-done:
				goto fin
			default:
				k = 2*k + 1
				if k == 1 {
					k++
				}
				out <- k
			}
		}
		for i := 0; i < q; i++ {
			ch, f = pg.genPrime(pg.bLimit, o)
			o += pg.bLimit
			for v := range ch {
				select {
				case <-done:
					f()
					goto fin
				default:
					out <- v
				}
			}
		}
		if r > 0 {
			ch, f = pg.genPrime(r, o)
			for v := range ch {
				select {
				case <-done:
					f()
					goto fin
				default:
					out <- v
				}
			}
		}
	fin:
		close(out)
	}()
	return out, func() { close(done) }
}
func (pg *primeGenerator) genPrime(n int, o int) (<-chan int, func()) {
	nn := n/2 + n%2
	if o%2 == 1 && n%2 == 1 {
		nn--
	}
	oo := o/2 + o%2
	cmpst := make([]bool, nn)
	l := int(math.Sqrt(float64(n + o)))
	c := cmpst
	if o != 0 {
		c = pg.m
	}
	for i := 3; i <= l; i += 2 {
		if c[(i-1)/2] {
			continue
		}
		j := i * i
		for j <= n+o {
			if j > o {
				cmpst[(j-1)/2-oo] = true
			}
			j += 2 * i
		}
	}
	ch := make(chan int)
	done := make(chan struct{})
	go func() {
		var v int
		for i := 0; i < len(cmpst); i++ {
			if !cmpst[i] {
				select {
				case <-done:
					break
				default:
					v = 2*i + 1 + o + o%2
					if v == 1 {
						v++
					}
					ch <- v
				}
			}
		}
		close(ch)
	}()
	return ch, func() { close(done) }
}
