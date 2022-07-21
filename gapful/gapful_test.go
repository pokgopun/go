/* https://theweeklychallenge.org/blog/perl-weekly-challenge-047/
TASK #2
Gapful Number
Write a script to print first 20 Gapful Numbers greater than or equal to 100. Please check out the page for more information about Gapful Numbers.
*/
package gapful

import "testing"

func TestIsGapful(t *testing.T) {
	g := []int{100, 105, 108, 110, 120, 121, 130, 132, 135, 140, 143, 150, 154, 160, 165, 170, 176, 180, 187, 190, 192, 195, 198, 200, 220, 225, 231, 240, 242, 253, 260, 264, 275, 280, 286, 297, 300, 315, 330, 341, 352, 360, 363, 374, 385, 390, 396, 400, 405, 440, 451}
	for i := 1; i < len(g); i += 2 {
		//fmt.Println("gapful =>", g[i-1:i+1])
		for _, j := range g[i-1 : i+1] {
			if IsGapful(j) == false {
				t.Errorf("incorrect result: expected true, got false")
			}
		}
		for j := g[i-1] + 1; j < g[i]; j++ {
			//fmt.Println("gapless =>", j)
			if IsGapful(j) == true {
				t.Errorf("incorrect result: expected false, got true")
			}
		}
	}
}
