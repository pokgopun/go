package main

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
)

func main() {
	fmt.Println(base35To10("1A"))
	fmt.Println(base35To10("11B"))
	fmt.Println(base10To35(365))
	fmt.Println(base10To35(366))
	fmt.Println(base10To35(330))
	fmt.Println(base10To35(18))
}

func byteToInt(b byte) int {
	if b < 58 {
		b -= 48
	} else {
		b -= 55
	}
	return int(b)
}

func intToByte(n int) byte {
	if n < 10 {
		n += 48
	} else {
		n += 55
	}
	return byte(n)
}

func base35To10(s string) (r int, err error) {
	if !regexp.MustCompile(`^[0-9A-Y]+$`).MatchString(s) {
		return 0, errors.New("Not agreed base35 strings")
	}
	last := len(s) - 1
	b := []byte(s)
	r += byteToInt(b[last])
	k := 1
	for last > 0 {
		last--
		k *= 35
		r += byteToInt(b[last]) * k
	}
	return r, nil
}
func base10To35(n int) (s string, err error) {
	bs := []byte{}
	b := 35
	for {
		r := n % b
		bs = append(bs, intToByte(r))
		if n = n / b; n == 0 {
			break
		}
	}
	sort.SliceStable(bs, func(i, j int) bool {
		return true
	})
	return string(bs), nil
}
