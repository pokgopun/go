/* https://theweeklychallenge.org/blog/perl-weekly-challenge-010/

Challenge #1
Write a script to encode/decode Roman numerals. For example, given Roman numeral CCXLVI, it should return 246. Similarly, for decimal number 39, it should return XXXIX. Checkout wikipedia page for more informaiton.

*/
package roman

import (
	"errors"
	"strings"
)

type Roman struct {
	keys []byte
	vals []uint
}

func NewRoman() Roman {
	return Roman{[]byte("IVXLCDM"), []uint{1, 5, 10, 50, 100, 500, 1000}}
}

func (r Roman) Digit(d, t uint) (b []byte, err error) {
	if d < 1 || d > 9 || t > 3 || (d > 3 && t > 2) {
		return []byte{}, errors.New("invalid digit or tenpower")
	}
	switch d {
	case 4, 9:
		b = append(b, r.keys[2*t], r.keys[2*t+1+d/5])
	case 5, 6, 7, 8:
		b = append(b, r.keys[2*t+1])
		fallthrough
	default:
		for i := uint(0); i < d%5; i++ {
			b = append(b, r.keys[2*t])
		}
	}
	return b, nil
}

func (r Roman) ToDec(str string) (n uint, err error) {
	if str != "" {
		var (
			b []byte
		)
		for t := 3; t >= 0; t-- {
			for d := 9; d > 0; d-- {
				b, err = r.Digit(uint(d), uint(t))
				if err == nil && strings.HasPrefix(str, string(b)) {
					var ten uint = 1
					for i := 0; i < t; i++ {
						ten *= 10
					}
					n += ten * uint(d)
					str = str[len(b):]
					if len(str) == 0 {
						return n, nil
					}
					break
				}
			}
		}
	}
	return 0, errors.New("not in roman standard form")
}

func (r Roman) FromDec(n uint) (string, error) {
	if n < 1 || n > 3999 {
		return "", errors.New("number must be greater than 0 and less than 4000")
	}
	var (
		d, t  uint
		b, bs []byte
		err   error
	)
	for n > 0 {
		d = n % 10
		if d > 0 {
			b, err = r.Digit(d, t)
			if err != nil {
				return "", err
			}
			bs = append(b, bs...)
		}
		n /= 10
		t++
	}
	return string(bs), nil
}
