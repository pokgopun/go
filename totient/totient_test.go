/* https://theweeklychallenge.org/blog/perl-weekly-challenge-047/
TASK #2
Gapful Number
Write a script to print first 20 Gapful Numbers greater than or equal to 100. Please check out the page for more information about Gapful Numbers.
*/
package totient

import "testing"

func BenchmarkIsPerfect(b *testing.B) {
	//perfect := []uint{3, 9, 15, 27, 39, 81, 111, 183, 243, 255, 327, 363, 471, 729, 2187, 2199, 3063, 4359, 4375, 5571}
	perfect := []uint{3, 9, 15, 27, 39, 81, 111, 183, 243, 255, 327, 363, 471, 729, 2187, 2199, 3063, 4359, 4375, 5571, 6561, 8751, 15723, 19683, 36759}
	processor := New(uint(len(perfect)))
	for k := 0; k < b.N; k++ {
		for i := 1; i < len(perfect); i += 2 {
			for _, j := range perfect[i-1 : i+1] {
				if processor.IsPerfect(j) == false {
					b.Errorf("incorrect result: expected true, got false")
				}
			}
			for j := perfect[i-1] + 1; j < perfect[i]; j++ {
				if processor.IsPerfect(j) == true {
					b.Errorf("incorrect result: expected false, got true")
				}
			}
		}
	}
}
