// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Simplify the Leap Year
//
//  1. Look at the solution of "the previous exercise".
//
//  2. And simplify the code (especially the if statements!).
//
// EXPECTED OUTPUT
//  It's the same as the previous exercise.
// ---------------------------------------------------------

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Give me a year number")
		return
	}
	if year, err := strconv.Atoi(args[1]); err != nil {
		fmt.Println("Error:", err)
	} else {
		if len(args[1]) != 4 {
			fmt.Printf("%q is not a valid year.\n", args[1])
		} else {

			if year%4 == 0 && (year%400 != 0 || year%4 == 0) {
				fmt.Printf("%d is a leap year.\n", year)
			} else {
				fmt.Printf("%d is not a leap year.\n", year)
			}
		}
	}
}
