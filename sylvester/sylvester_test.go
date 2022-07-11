/* https://theweeklychallenge.org/blog/perl-weekly-challenge-173/
Task 2: Sylvester’s sequence
Submitted by: Mohammad S Anwar
Write a script to generate first 10 members of Sylvester's sequence. For more informations, please refer to the wikipedia page.


Output
2
3
7
43
1807
3263443
10650056950807
113423713055421844361000443
12864938683278671740537145998360961546653259485195807
165506647324519964198468195444439180017513152706377497841851388766535868639572406808911988131737645185443
*/
package sylvester

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Sylvester(t *testing.T) {
	sample := "2, 3, 7, 43, 1807, 3263443, 10650056950807, 113423713055421844361000443, 12864938683278671740537145998360961546653259485195807, 165506647324519964198468195444439180017513152706377497841851388766535868639572406808911988131737645185443"
	count := uint(len(strings.Split(sample, ", ")))
	expected := strings.ReplaceAll(sample, ", ", "\n")
	if diff := cmp.Diff(expected, Sylvester(count)); diff != "" {
		t.Error(diff)
	}
}
