package roman

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDigit(t *testing.T) {
	data := []struct {
		t, d        uint
		res, errMsg string
	}{
		{0, 1, "I", ""},
		{0, 2, "II", ""},
		{0, 3, "III", ""},
		{0, 4, "IV", ""},
		{0, 5, "V", ""},
		{0, 6, "VI", ""},
		{0, 7, "VII", ""},
		{0, 8, "VIII", ""},
		{0, 9, "IX", ""},
		{1, 1, "X", ""},
		{1, 2, "XX", ""},
		{1, 3, "XXX", ""},
		{1, 4, "XL", ""},
		{1, 5, "L", ""},
		{1, 6, "LX", ""},
		{1, 7, "LXX", ""},
		{1, 8, "LXXX", ""},
		{1, 9, "XC", ""},
		{2, 1, "C", ""},
		{2, 2, "CC", ""},
		{2, 3, "CCC", ""},
		{2, 4, "CD", ""},
		{2, 5, "D", ""},
		{2, 6, "DC", ""},
		{2, 7, "DCC", ""},
		{2, 8, "DCCC", ""},
		{2, 9, "CM", ""},
		{3, 1, "M", ""},
		{3, 2, "MM", ""},
		{3, 3, "MMM", ""},
		{3, 4, "", "invalid digit or tenpower"},
		{0, 0, "", "invalid digit or tenpower"},
		{1, 10, "", "invalid digit or tenpower"},
		{4, 1, "", "invalid digit or tenpower"},
	}
	roman := NewRoman()
	for _, d := range data {
		res, err := roman.Digit(d.d, d.t)
		if diff := cmp.Diff(d.res, string(res)); diff != "" {
			t.Error(diff)
		}
		var errMsg string
		if err != nil {
			errMsg = err.Error()
		}
		if errMsg != d.errMsg {
			t.Errorf("incorrect result: expected `%s`, got `%s`", d.errMsg, errMsg)
		}
	}
}

func TestFromDec(t *testing.T) {
	data := []struct {
		n           uint
		res, errMsg string
	}{
		{246, "CCXLVI", ""},
		{39, "XXXIX", ""},
		{0, "", "number must be greater than 0 and less than 4000"},
		{4000, "", "number must be greater than 0 and less than 4000"},
		{2022, "MMXXII", ""},
	}
	roman := NewRoman()
	for _, d := range data {
		res, err := roman.FromDec(d.n)
		if res != d.res {
			t.Errorf("incorrect result: expected `%s`, got `%s`", d.res, res)
		}
		var errMsg string
		if err != nil {
			errMsg = err.Error()
		}
		if errMsg != d.errMsg {
			t.Errorf("incorrect result: expected `%s`, got `%s`", d.errMsg, errMsg)
		}
	}
}

func TestToDec(t *testing.T) {
	data := []struct {
		res         uint
		str, errMsg string
	}{
		{246, "CCXLVI", ""},
		{39, "XXXIX", ""},
		{0, "", "not in roman standard form"},
		{0, "MMMM", "not in roman standard form"},
		{0, "CDD", "not in roman standard form"},
		{0, "XXC", "not in roman standard form"},
		{0, "VIIII", "not in roman standard form"},
		{2022, "MMXXII", ""},
	}
	roman := NewRoman()
	for _, d := range data {
		res, err := roman.ToDec(d.str)
		if res != d.res {
			t.Errorf("incorrect result: expected `%d`, got `%d`", d.res, res)
		}
		var errMsg string
		if err != nil {
			errMsg = err.Error()
		}
		if errMsg != d.errMsg {
			t.Errorf("incorrect result: expected `%s`, got `%s`", d.errMsg, errMsg)
		}
	}
}
