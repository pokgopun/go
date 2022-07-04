/* https://theweeklychallenge.org/blog/perl-weekly-challenge-166/
Task 1: Hexadecimal Words

Submitted by: [43]Ryan J Thompson
     __________________________________________________________________

   As an old systems programmer, whenever I needed to come up with a
   32-bit number, I would reach for the tired old examples like 0xDeadBeef
   and 0xC0dedBad. I want more!

   Write a program that will read from a dictionary and find 2- to
   8-letter words that can be “spelled” in hexadecimal, with the addition
   of the following letter substitutions:
     * o ⟶ 0 (e.g., 0xf00d = “food”)
     * l ⟶ 1
     * i ⟶ 1
     * s ⟶ 5
     * t ⟶ 7

   You can use your own dictionary or you can simply open
   ../../../data/dictionary.txt (relative to your script’s location in our
   [44]GitHub repository) to access the dictionary of common words from
   [45]Week #161.

Optional Extras (for an 0xAddedFee, of course!)

    1. Limit the number of “special” letter substitutions in any one
       result to keep that result at least somewhat comprehensible.
       (0x51105010 is an actual example from my sample solution you may
       wish to avoid!)
    2. Find phrases of words that total 8 characters in length (e.g.,
       0xFee1Face), rather than just individual words.

*/
package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	var ds struct {
		wib  string
		wibs []string
	}
	//ds.wibs = make([][]byte, 40)
	f, err := os.Open("../../../data/dictionary.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	//scanner.Split(bufio.ScanWords)
	re := regexp.MustCompile(`^[abcdefolist]{2,8}$`)
	if err != nil {
		log.Fatal(err)
	}
	for scanner.Scan() {
		ds.wib = scanner.Text()
		if re.MatchString(ds.wib) {
			ds.wibs = append(ds.wibs, ds.wib)
		}
	}
	re = regexp.MustCompile(`[abcdef]`)
	sort.SliceStable(ds.wibs, func(i, j int) bool {
		lenOddI := len(re.ReplaceAllString(ds.wibs[i], ""))
		lenOddJ := len(re.ReplaceAllString(ds.wibs[j], ""))
		if lenOddI == lenOddJ {
			return len(ds.wibs[i]) > len(ds.wibs[j])
		} else {
			return lenOddI < lenOddJ
		}
	})
	r := strings.NewReplacer("o", "0", "l", "1", "i", "1", "s", "5", "t", "7")
	w := bufio.NewWriter(os.Stdout)
	for _, v := range ds.wibs {
		w.WriteString(v + " => 0x" + r.Replace(v) + "\n")
	}
	w.Flush()
}
