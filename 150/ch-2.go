/* https://theweeklychallenge.org/blog/perl-weekly-challenge-150/

TASK #2 › Square-free Integer

Submitted by: [41]Mohammad S Anwar
     __________________________________________________________________

   Write a script to generate all square-free integers <= 500.

     In mathematics, a square-free integer (or squarefree integer) is an
     integer which is divisible by no perfect square other than 1. That
     is, its prime factorization has exactly one factor for each prime
     that appears in it. For example, 10 = 2 ⋅ 5 is square-free, but 18 =
     2 ⋅ 3 ⋅ 3 is not, because 18 is divisible by 9 = 3**2.

Example

The smallest positive square-free integers are
    1, 2, 3, 5, 6, 7, 10, 11, 13, 14, 15, 17, 19, 21, 22, 23, 26, 29, 30, ...
*/
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	l := 500
	if len(os.Args) > 1 {
		i, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		l = i
	}
	fmt.Println("The smallest positive square-free integers are")
	sf := []uint{1}
	//fmt.Print("\t1")
	chk, nxt, _ := primeUtil()
	for i := 2; i <= l; i++ {
		if chk(i) {
			sf = append(sf, uint(i))
			//fmt.Print(", ", i)
		} else {
			n := i
			p := 2
			seen := 0
			for {
				if n%p != 0 {
					p = nxt(p)
				} else {
					if p == seen {
						break
					} else {
						seen = p
					}
					n /= p
					if n == 1 {
						sf = append(sf, uint(i))
						//fmt.Print(", ", i)
						break
					} else if chk(n) {
						if n != seen {
							sf = append(sf, uint(i))
							//fmt.Print(", ", i)
						}
						break
					}
				}
			}
		}

	}
	fmt.Println(sf)
	//fmt.Print("\n")
}

func primeUtil() (func(i int) bool, func(i int) int, func()) {
	m := map[int]bool{}
	return func(i int) bool {
			is, ok := m[i]
			if ok {
				//fmt.Println("### check prime from cache ###")
				return is
			} else {
				if isPrime(i) {
					m[i] = true
					return true
				}
				m[i] = false
				return false
			}
		}, func(i int) int {
			for {
				i++
				is, ok := m[i]
				if ok {
					//fmt.Println("### check next number from cache ###")
					if is {
						return i
					} else {
						continue
					}
				} else {
					if isPrime(i) {
						m[i] = true
						return i
					}
					m[i] = false
				}
			}
		}, func() {
			fmt.Println("=>", m)
		}
}
func isPrime(n int) bool {
	//fmt.Println("In isPrime")
	if n <= 3 {
		return n > 1
	} else if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
