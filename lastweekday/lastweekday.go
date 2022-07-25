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
	"fmt"
	"time"
)

func monthdayCount(y uint) (r [12]int) {
	r = [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if y%4 == 0 {
		r[time.February]++
	}
	return r
}

type lastWeekdayOfMonth struct {
	days   [12]time.Time
	format string
}

func (lwdom lastWeekdayOfMonth) String() (r string) {
	for _, t := range lwdom.days {
		r += fmt.Sprintln(t.Format(lwdom.format))
	}
	return r
}
func New(wd time.Weekday, y uint, format string) (lwdom lastWeekdayOfMonth) {
	lwdom.format = format
	for i, v := range monthdayCount(y) {
		// Find offset between the last day of the month and the specified weekday
		t := time.Date(int(y), time.Month(i+1), v, 0, 0, 0, 0, time.UTC)
		o := int(wd - t.Weekday())
		// The offset is the number of day we need to go back to the latest specified weekday in certain month
		if o != 0 {
			// plus sign indicate the offset between specified weekday and the current weekday of the previous week
			if o > 0 {
				o -= 7
			}
		}
		lwdom.days[i] = t.AddDate(0, 0, o)
	}
	return lwdom
}
