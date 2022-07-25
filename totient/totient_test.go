/* https://theweeklychallenge.org/blog/perl-weekly-challenge-047/
TASK #2
Gapful Number
Write a script to print first 20 Gapful Numbers greater than or equal to 100. Please check out the page for more information about Gapful Numbers.
*/
package totient

import "testing"

func TestIsPerfect(t *testing.T) {
	perfect := []uint{3, 9, 15, 27, 39, 81, 111, 183, 243, 255, 327, 363, 471, 729, 2187, 2199, 3063, 4359, 4375, 5571}
	processor := New()
	for i := 1; i < len(perfect); i += 2 {
		for _, j := range perfect[i-1 : i+1] {
			if processor.IsPerfect(j) == false {
				t.Errorf("incorrect result: expected true, got false")
			}
		}
		for j := perfect[i-1] + 1; j < perfect[i]; j++ {
			if processor.IsPerfect(j) == true {
				t.Errorf("incorrect result: expected false, got true")
			}
		}
	}
}
