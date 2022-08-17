/* https://theweeklychallenge.org/blog/perl-weekly-challenge-138/
TASK #2 › Split Number
Submitted by: Mohammad S Anwar
You are given a perfect square.

Write a script to figure out if the square root the given number is same as sum of 2 or more splits of the given number.

Example 1
Input: $n = 81
Output: 1

Since, sqrt(81) = 8 + 1
Example 2
Input: $n = 9801
Output: 1

Since, sqrt(9801) = 98 + 0 + 1
Example 3
Input: $n = 36
Output: 0

Since, sqrt(36) != 3 + 6
*/
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	n := 9801
	fmt.Sscanf(strings.Join(os.Args[1:], " "), "%d", &n)
	fmt.Printf("Input: n = %d\nOutput: %t\n", n, isSplitNum(n))
}
func isSplitNum(n int) bool {
	sqrt := int(math.Sqrt(float64(n)))
	if n != sqrt*sqrt {
		return false
	}
	nStr := strconv.Itoa(n)
	var res string
	splits := strings.TrimSuffix(numSplit(nStr, "", &res), " "+nStr)
	for _, v := range strings.Split(splits, " ") {
		if sumSplit(v) == sqrt {
			//fmt.Printf("Since, sqrt(%d) = %d = %s\n", n, sqrt, v)
			return true
		}
	}
	return false
}

func numSplit(s, t string, res *string) string {
	l := len(s)
	if l > 0 {
		for i := 0; i < l; i++ {
			numSplit(s[i+1:], strings.Join([]string{t, s[:i+1]}, "+"), res)
		}
	} else {
		*res += " " + t[1:]
		return ""
	}
	return (*res)[1:]
}

func sumSplit(str string) (sum int) {
	for _, v := range strings.Split(str, "+") {
		n, _ := strconv.Atoi(v)
		sum += n
	}
	return sum
}
