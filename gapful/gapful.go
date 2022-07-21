/* https://theweeklychallenge.org/blog/perl-weekly-challenge-047/
TASK #2
Gapful Number
Write a script to print first 20 Gapful Numbers greater than or equal to 100. Please check out the page for more information about Gapful Numbers.
*/
package gapful

func IsGapful(i int) bool {
	n := i
	var b []byte
	for n > 0 {
		b = append(b, byte(n%10))
		n /= 10
	}
	d := int(b[0] + 10*b[len(b)-1])
	if i%d == 0 {
		return true
	}
	return false
}
