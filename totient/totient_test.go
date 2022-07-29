/* https://theweeklychallenge.org/blog/perl-weekly-challenge-047/
TASK #2
Gapful Number
Write a script to print first 20 Gapful Numbers greater than or equal to 100. Please check out the page for more information about Gapful Numbers.
*/
package totient

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func BenchmarkIsPerfect(b *testing.B) {
	perfect := []uint64{3, 9, 15, 27, 39, 81, 111, 183, 243, 255, 327, 363, 471, 729, 2187, 2199, 3063, 4359, 4375, 5571}
	//perfect := []uint64{3, 9, 15, 27, 39, 81, 111, 183, 243, 255, 327, 363, 471, 729, 2187, 2199, 3063, 4359, 4375, 5571, 6561, 8751, 15723, 19683, 36759, 46791}
	l := len(perfect)
	processor := New()
	for k := 0; k < b.N; k++ {
		i := 0
		if processor.IsPerfect(perfect[i]) == false {
			b.Errorf("IsPerfect(%d): incorrect result: expected true, got false", perfect[i])
		}
		for i = 1; i < l; i += 1 {
			for j := perfect[i-1] + 1; j < perfect[i]; j++ {
				if processor.IsPerfect(j) == true {
					b.Errorf("perfect(%d): incorrect result: expected false, got true", j)
				}
			}
			if processor.IsPerfect(perfect[i]) == false {
				b.Errorf("IsPerfect(%d): incorrect result: expected true, got false", perfect[i])
			}
		}
	}
}

func BenchmarkListPerfect(b *testing.B) {
	data := []struct {
		start, count uint64
		res          []uint64
	}{
		/**/
		{
			3, 20,
			[]uint64{
				3, 9, 15, 27, 39,
				81, 111, 183, 243, 255,
				327, 363, 471, 729, 2187,
				2199, 3063, 4359, 4375, 5571,
			},
		},
		/**/
		/*
			{
				3, 26,
				[]uint64{
					3, 9, 15, 27, 39,
					81, 111, 183, 243, 255,
					327, 363, 471, 729, 2187,
					2199, 3063, 4359, 4375, 5571,
					6561, 8751, 15723, 19683, 36759, 46791,
				},
			},
		*/
	}
	processor := New()
	for k := 0; k < b.N; k++ {
		for _, d := range data {
			if diff := cmp.Diff(d.res, processor.ListPerfect(d.start, d.count)); diff != "" {
				b.Error(diff)
			}
		}
	}
}
