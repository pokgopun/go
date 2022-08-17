/* https://theweeklychallenge.org/blog/perl-weekly-challenge-137/
TASK #1 › Long Year
Submitted by: Mohammad S Anwar
Write a script to find all the years between 1900 and 2100 which is a Long Year.

A year is Long if it has 53 weeks.


[UPDATED][2021-11-01 16:20:00]: For more information about Long Year, please refer to wikipedia.

Expected Output
1903, 1908, 1914, 1920, 1925,
1931, 1936, 1942, 1948, 1953,
1959, 1964, 1970, 1976, 1981,
1987, 1992, 1998, 2004, 2009,
2015, 2020, 2026, 2032, 2037,
2043, 2048, 2054, 2060, 2065,
2071, 2076, 2082, 2088, 2093,
2099
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	startY, endY := 1900, 2100
	var longYear []int
	for y := startY; y <= endY; y++ {
		t := time.Date(y, time.January, 1, 0, 0, 0, 0, time.UTC)
		for i := 0; i <= 1; i++ {
			if t.AddDate(i, 0, -i).Weekday() == time.Thursday {
				longYear = append(longYear, y)
				break
			}
		}
	}
	fmt.Println(longYear)
}
