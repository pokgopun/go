/* https://theweeklychallenge.org/blog/perl-weekly-challenge-149/
TASK #1 › Fibonacci Digit Sum

Submitted by: [38]Roger Bell_West
     __________________________________________________________________

   Given an input $N, generate the first $N numbers for which the sum of
   their digits is a Fibonacci number.

Example

f(20)=[0, 1, 2, 3, 5, 8, 10, 11, 12, 14, 17, 20, 21, 23, 26, 30, 32, 35, 41, 44]

*/
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var n int
	if len(os.Args) > 1 {
		i, err := strconv.ParseUint(os.Args[1], 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		n = int(i)
	} else {
		n = 20
	}
	fds := make([]int, n)
	i := 0
	//isFib, t := makeIsFib()
	isFib, _ := makeIsFib()
	for j := 0; i < n; j++ {
		var sum int
		if j < 10 {
			sum = j
		} else {
			for _, v := range []byte(strconv.Itoa(j)) {
				k, err := strconv.Atoi(string(v))
				if err != nil {
					log.Fatal(err)
				}
				sum += k
			}
		}
		if isFib(sum) {
			fds[i] = j
			i++
		}
	}
	fmt.Println(fds)
	//t()
}

func makeIsFib() (func(int) bool, func()) {
	m := map[int]bool{
		0: true,
		1: true,
	}
	a, b := 1, 1
	return func(n int) bool {
			for b < n {
				a, b = b, b+a
				m[b] = true
			}
			return m[n]

		}, func() {
			fmt.Println(m)
		}
}
