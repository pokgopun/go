/* https://theweeklychallenge.org/blog/perl-weekly-challenge-164/

Task 2: Happy Numbers

Submitted by: [56]Robert DiCicco
     __________________________________________________________________

   Write a script to find the first 8 Happy Numbers in base 10. For more
   information, please check out [57]Wikipedia.

   Starting with any positive integer, replace the number by the sum of
   the squares of its digits, and repeat the process until the number
   equals 1 (where it will stay), or it loops endlessly in a cycle which
   does not include 1.

   Those numbers for which this process end in 1 are happy numbers, while
   those numbers that do not end in 1 are unhappy numbers.

Example

   19 is Happy Number in base 10, as shown:
19 => 1^2 + 9^2
   => 1   + 81
   => 82 => 8^2 + 2^2
         => 64  + 4
         => 68 => 6^2 + 8^2
               => 36  + 64
               => 100 => 1^2 + 0^2 + 0^2
                      => 1 + 0 + 0
                      => 1
*/
package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	//n := 10_000_000
	n := 8
	l := 10_000_000
	//l := 1000
	if len(os.Args) > 1 {
		r, err := strconv.Atoi(os.Args[1])
		if err == nil {
			n = r
		}
	}
	w := bufio.NewWriter(os.Stdout)
	w.WriteString("1")
	count := 1
	var h happy = make(map[int]bool)
	for i := 2; count < n && i <= l; i++ {
		if h.isHappy(i, []int{}) {
			w.WriteString(", " + strconv.Itoa(i))
			count++
		}
	}
	w.WriteString("\n")
	w.Flush()
	//fmt.Println(len(happy), happy)
}

type happy map[int]bool

func (h *happy) isHappy(i int, ht []int) bool {
	happy := *h
	b, ok := happy[i]
	if ok {
		return b
	}
	for _, v := range ht {
		if i == v {
			for _, v := range ht {
				happy[v] = false
			}
			return false
		}
	}
	var sum int
	//fmt.Print(i, " => ")
	for _, v := range strconv.Itoa(i) {
		v -= 48
		sum += int(v * v)
		//fmt.Printf("%v^2 ", v)
	}
	//fmt.Println("=>", sum)
	ht = append(ht, i)
	if sum == 1 {
		for _, v := range ht {
			happy[v] = true
		}
		return true
	}
	return h.isHappy(sum, ht)
}
