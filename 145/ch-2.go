/* https://theweeklychallenge.org/blog/perl-weekly-challenge-145/

TASK #2 › Palindromic Tree

Submitted by: [46]Mohammad S Anwar
     __________________________________________________________________

   You are given a string $s.

   Write a script to create a Palindromic Tree for the given string.

   I found this [47]blog exaplaining Palindromic Tree in detail.

Example 1:

Input: $s = 'redivider'
Output: r redivider e edivide d divid i ivi v

Example 2:

Input: $s = 'deific'
Output: d e i ifi f c

Example 3:

Input: $s = 'rotors'
Output: r rotor o oto t s

Example 4:

Input: $s = 'challenge'
Output: c h a l ll e n g

Example 5:

Input: $s = 'champion'
Output: c h a m p i o n

Example 6:

Input: $s = 'christmas'
Output: c h r i s t m a
*/
// All examples do not have a case that a letter happens more than twice in given words but we will handle the case as well
// Here additional examples for the case: banana, redeemable, deterministic, heterogeneity, initiation, monopolous, honolulu
package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var sample []string
	if len(os.Args) > 1 {
		sample = os.Args[1:]
	} else {
		sample = []string{
			"redivider",
			"deific",
			"rotors",
			"challenge",
			"champion",
			"christmas",
			"banana",
			"redeemable",
			"deterministic",
			"heterogeneity",
			"initiation",
			"monopolous",
			"honolulu",
		}
	}
	for _, v := range sample {
		p := newPldt(v)
		fmt.Println(p)
	}
}

type pldt struct {
	word string
	seen map[string]bool
	cpos map[byte][]int
	vals []string
}

func newPldt(s string) (p pldt) {
	p.word = s
	p.seen = make(map[string]bool)
	p.cpos = make(map[byte][]int)
	for i, v := range []byte(p.word) {
		p.cpos[v] = append(p.cpos[v], i)
	}
	for _, c := range []byte(p.word) {
		if p.seen[string(c)] {
			continue
		}
		p.seen[string(c)] = true
		p.vals = append(p.vals, string(c))
		for i, pos1 := range p.cpos[c] {
			for _, pos2 := range p.cpos[c][i+1:] {
				o := p.word[pos1 : pos2+1]
				if p.seen[o] {
					continue
				}
				r := []byte(o)
				sort.SliceStable(r, func(i, j int) bool {
					return true
				})
				if o == string(r) {
					p.vals = append(p.vals, o)
					p.seen[o] = true
				}
			}
		}
	}
	return p
}

func (p pldt) String() string {
	var b strings.Builder
	b.WriteString("Input: '" + p.word + "'\nOutput: " + strings.Join(p.vals, " ") + "\n")
	return b.String()
}
