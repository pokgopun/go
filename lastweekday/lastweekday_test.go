/*
https://theweeklychallenge.org/blog/perl-weekly-challenge-013/
Challenge #1

     Write a script to print the date of last Friday of every month of a
     given year. For example, if the given year is 2019 then it should
     print the following:

2019/01/25
2019/02/22
2019/03/29
2019/04/26
2019/05/31
2019/06/28
2019/07/26
2019/08/30
2019/09/27
2019/10/25
2019/11/29
2019/12/27

https://theweeklychallenge.org/blog/perl-weekly-challenge-175/
Task 1: Last Sunday
Submitted by: Mohammad S Anwar
Write a script to list Last Sunday of every month in the given year.

For example, for year 2022, we should get the following:


2022-01-30
2022-02-27
2022-03-27
2022-04-24
2022-05-29
2022-06-26
2022-07-31
2022-08-28
2022-09-25
2022-10-30
2022-11-27
2022-12-25
*/
package lastweekday

import (
	"testing"
	"time"
)

func TestString(t *testing.T) {
	data := []struct {
		weekday     time.Weekday
		year        uint
		format, res string
	}{
		{time.Friday, 2019, "2006/01/02", `2019/01/25
2019/02/22
2019/03/29
2019/04/26
2019/05/31
2019/06/28
2019/07/26
2019/08/30
2019/09/27
2019/10/25
2019/11/29
2019/12/27
`},
		{time.Sunday, 2022, "2006-01-02", `2022-01-30
2022-02-27
2022-03-27
2022-04-24
2022-05-29
2022-06-26
2022-07-31
2022-08-28
2022-09-25
2022-10-30
2022-11-27
2022-12-25
`},
	}
	for _, d := range data {
		if res := New(d.weekday, d.year, d.format).String(); res != d.res {
			t.Errorf("incorrect result, expected %s, got %s", d.res, res)
		}
	}
}
