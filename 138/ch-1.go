/* https://theweeklychallenge.org/blog/perl-weekly-challenge-138/
TASK #1 › Workdays
Submitted by: Mohammad S Anwar
You are given a year, $year in 4-digits form.

Write a script to calculate the total number of workdays in the given year.

For the task, we consider, Monday - Friday as workdays.

Example 1
Input: $year = 2021
Output: 261
Example 2
Input: $year = 2020
Output: 262
*/
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	year := 2021
	fmt.Sscanf(strings.Join(os.Args[1:], " "), "%d", &year)
	workdays := 52 * 5 // every year has at least 52 weeks each of which has 5 work days
	daysLeft := 1      // non-leap year has 365 - (52 - 7) = 1 day left to check if it is a working day
	if year%4 == 0 {   // a leap year will add another day to daysLeft
		daysLeft++
	}
	// check if there a workday in daysLeft, start at the end of the year
	t := time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC)
	for d := 0; d < daysLeft; d++ {
		if wd := t.AddDate(0, 0, -d).Weekday(); wd > time.Sunday && wd < time.Saturday {
			workdays++
		}
	}
	fmt.Printf("Input: year = %d\nOutput: %d\n", year, workdays)
}
