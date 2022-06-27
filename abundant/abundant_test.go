package abundant

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const datafile = "testdata/b005101.txt"

func TestOddAbundant(t *testing.T) {
	f, err := os.Open(datafile)
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	var (
		n, a uint
		s    []uint
	)
	for {
		_, err = fmt.Fscanln(r, &n, &a)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if a%2 != 0 {
			s = append(s, a)
		}
	}
	if diff := cmp.Diff(s, OddAbundant(uint(len(s)))); diff != "" {
		t.Error(diff)
	}
}
func TestIsAbundant(t *testing.T) {
	f, err := os.Open(datafile)
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	var (
		i, n, a  uint
		answer   []bool
		ans, res bool
	)
	for {
		_, err = fmt.Fscanln(r, &n, &a)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		answer = make([]bool, a-i+1)
		answer[a-i] = true
		for j := i; j <= a; j++ {
			res, ans = IsAbundant(j), answer[j-i]
			//fmt.Println(j, "=>", res)
			if res != ans {
				t.Errorf("%d: expected %t, got %t", j, ans, res)
			}
		}
		i = a + 1
	}
}
