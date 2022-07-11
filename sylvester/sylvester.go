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
	"math/big"
	"strings"
)

func Sylvester(count uint) string {
	if count == 0 {
		return ""
	}
	n := big.NewInt(2)
	var b strings.Builder
	for count > 0 {
		b.WriteString("\n" + n.String())
		n.Add(new(big.Int).Mul(n, n), new(big.Int).Sub(big.NewInt(1), n))
		count--
	}
	return b.String()[1:]
}
