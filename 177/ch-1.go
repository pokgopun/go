/* https://theweeklychallenge.org/blog/perl-weekly-challenge-177/
Task 1: Damm Algorithm
Submitted by: Mohammad S Anwar
You are given a positive number, $n.

Write a script to validate the given number against the included check digit.

Please checkout the wikipedia page for information.

Example 1
Input: $n = 5724
Output: 1 as it is valid number
Example 2
Input: $n = 5727
Output: 0 as it is invalid number
*/
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	tStr := "0317598642709215486342068713591750983426612304597836742095815869720134894536201794386172052581436790"
	guide := "please provide a positive number as an argument"
	var n uint = 5724
	if len(os.Args) < 2 {
		log.Fatal(guide)
	}
	_, err := fmt.Sscanf(os.Args[1], "%d", &n)
	if err != nil {
		log.Fatal(guide)
	}
	var r byte
	for _, v := range []byte(os.Args[1]) {
		//fmt.Println("r =", r)
		v -= 48
		r = tStr[r*10+v] - 48
	}
	fmt.Printf("Input: n = %s\nOutput: %t\n", os.Args[1], r == 0)
}
