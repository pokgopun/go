/* https://theweeklychallenge.org/blog/perl-weekly-challenge-013/

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
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(newLastWeekdayOfMonth(time.Friday, 2019))
}
func monthdayCount(y int) (r [12]int) {
	r = [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if y%4 == 0 {
		r[time.February]++
	}
	return r
}

type lastWeekdayOfMonth [12]time.Time

func (lwdom lastWeekdayOfMonth) String() (r string) {
	for _, t := range lwdom {
		r += fmt.Sprintln(t.Format("2006/01/02"))
	}
	return r
}
func newLastWeekdayOfMonth(wd time.Weekday, y int) (lwdom lastWeekdayOfMonth) {
	for i, v := range monthdayCount(y) {
		// Find offset between the last day of the month and the specified weekday
		t := time.Date(y, time.Month(i+1), v, 0, 0, 0, 0, time.UTC)
		o := int(wd - t.Weekday())
		// The offset is the number of day we need to go back to the latest specified weekday in certain month
		if o != 0 {
			// plus sign indicate the offset between specified weekday and the current weekday of the previous week
			if o > 0 {
				o -= 7
			}
		}
		lwdom[i] = t.AddDate(0, 0, o)
	}
	return lwdom
}
