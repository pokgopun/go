/* https://theweeklychallenge.org/blog/perl-weekly-challenge-161/
 Task 1: Abecedarian Words

Submitted by: [53]Ryan J Thompson
     __________________________________________________________________

   An abecedarian word is a word whose letters are arranged in
   alphabetical order. For example, “knotty” is an abecedarian word, but
   “knots” is not. Output or return a list of all abecedarian words in the
   [54]dictionary, sorted in decreasing order of length.

   Optionally, using only abecedarian words, leave a short comment in your
   code to make your reviewer smile.

*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func main() {
	dict := "../../../data/dictionary.txt"
	if len(os.Args) > 1 {
		dict = os.Args[1]
	}
	var r io.Reader
	f, err := os.Open(dict)
	if err != nil {
		//log.Fatal(err)
		s := `forty
hippy
bees
buzz
a
dirty
nosy
chimp
`
		r = strings.NewReader(s)
	} else {
		defer f.Close()
		r = f
	}
	abc := []string{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s := scanner.Text()
		if isAbc(s) {
			abc = append(abc, s)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	_, ok := r.(*os.File)
	if ok {
		sort.SliceStable(abc, func(i, j int) bool {
			return len(abc[i]) > len(abc[j])
		})
	}
	fmt.Println(abc)
}
func isAbc(s string) bool {
	b := []byte(s)
	if len(b) > 1 {
		for i := 1; i < len(b); i++ {
			if b[i] < b[i-1] {
				return false
			}
		}
	}
	return true
}
