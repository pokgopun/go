/* https://theweeklychallenge.org/blog/perl-weekly-challenge-161/

Task 2: Pangrams

Submitted by: [55]Ryan J Thompson
     __________________________________________________________________

   A pangram is a sentence or phrase that uses every letter in the English
   alphabet at least once. For example, perhaps the most well known
   pangram is:
the quick brown fox jumps over the lazy dog

   Using the provided [56]dictionary, so that you don’t need to include
   individual copy, generate at least one pangram.

   Your pangram does not have to be a syntactically valid English sentence
   (doing so would require far more work, and a dictionary of nouns,
   verbs, adjectives, adverbs, and conjunctions). Also note that repeated
   letters, and even repeated words, are permitted.

BONUS: Constrain or optimize for something interesting (completely up to
you), such as:

Shortest possible pangram (difficult)
Pangram which contains only abecedarian words (see challenge 1)
Pangram such that each word "solves" exactly one new letter. For example, such a
 pangram might begin with (newly solved letters in bold):
    a ah hi hid die ice tea ...
    What is the longest possible pangram generated with this method? (All soluti
ons will contain 26 words, so focus on the letter count.)
Pangrams that have the weirdest (PG-13) Google image search results
Anything interesting goes!
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

type byteSeen map[byte]struct{}

func (bs byteSeen) countUnseen(s string) (n int) {
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	var seen byte
	for _, v := range b {
		if seen == v {
			continue
		} else {
			seen = v
		}
		_, ok := bs[v]
		if !ok {
			n++
		}
	}
	return n
}
func main() {
	dict := "../../../data/dictionary.txt"
	if len(os.Args) > 1 {
		dict = os.Args[1]
	}
	var r io.Reader
	f, err := os.Open(dict)
	if err != nil {
		//log.Fatal(err)
		s := `the
quick
brown
fox
jumps
over
the
lazy
dog
`
		r = strings.NewReader(s)
	} else {
		defer f.Close()
		r = f
	}
	pngrm := []string{}
	letters := byteSeen{}
	var best string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		b := scanner.Text()
		if len(best) == 0 {
			//best = "a"
			best = b[0:1]
		}
		if b[:1] == best[:1] {
			if letters.countUnseen(best) >= letters.countUnseen(b) {
				continue
			}
			best = b
			continue
		}
		b, best = best, b
		//fmt.Println("b=", b, "best=", best)
		for _, v := range []byte(b) {
			letters[v] = struct{}{}
		}
		pngrm = append(pngrm, b)
		//fmt.Println(len(letters), len(pngrm))
		if len(letters) == 26 {
			break
		}
	}
	if len(letters) < 26 {
		pngrm = append(pngrm, best)
		for _, v := range []byte(best) {
			letters[v] = struct{}{}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	if len(letters) == 26 {
		fmt.Println(pngrm)
		fmt.Printf("=> %d words, %d letters and %d unique letters\n", len(pngrm), len(strings.Join(pngrm, "")), len(letters))
	}
}
