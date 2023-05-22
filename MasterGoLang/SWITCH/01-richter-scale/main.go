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
// STORY
//  You're curious about the richter scales. When reporters
//  say: "There's been a 5.5 magnitude earthquake",
//
//  You want to know what that means.
//
//  So, you've decided to write a program to do that for you.
//
// EXERCISE: Richter Scale
//
//  1. Get the earthquake magnitude from the command-line.
//  2. Display its corresponding description.
//
// ---------------------------------------------------------
// MAGNITUDE                    DESCRIPTION
// ---------------------------------------------------------
// Less than 2.0                micro
// 2.0 and less than 3.0        very minor
// 3.0 and less than 4.0        minor
// 4.0 and less than 5.0        light
// 5.0 and less than 6.0        moderate
// 6.0 and less than 7.0        strong
// 7.0 and less than 8.0        major
// 8.0 and less than 10.0       great
// 10.0 or more                 massive
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me the magnitude of the earthquake.
//
//  go run main.go ABC
//    I couldn't get that, sorry.
//
//  go run main.go 0.5
//    0.50 is micro
//
//  go run main.go 2.5
//    2.50 is very minor
//
//  go run main.go 3
//    3.00 is minor
//
//  go run main.go 4.5
//    4.50 is light
//
//  go run main.go 5
//    5.00 is moderate
//
//  go run main.go 6
//    6.00 is strong
//
//  go run main.go 7
//    7.00 is major
//
//  go run main.go 8
//    8.00 is great
//
//  go run main.go 11
//    11.00 is massive
// ---------------------------------------------------------

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Give me the magnitude of the earthquake.")
		return
	}
	input := args[1]
	var magnitude string
	richter, err := strconv.ParseFloat(input, 32)
	if err != nil {
		fmt.Println("I couldn't get that, sorry.")
		return
	}
	switch richter > 0 {
	case richter < 2:
		magnitude = "micro"
	case richter >= 2 && richter < 3:
		magnitude = "very minor"
	case richter >= 3 && richter < 4:
		magnitude = "minor"
	case richter >= 4 && richter < 5:
		magnitude = "light"
	case richter >= 5 && richter < 6:
		magnitude = "moderate"
	case richter >= 6 && richter < 7:
		magnitude = "strong"
	case richter >= 7 && richter < 8:
		magnitude = "major"
	case richter >= 8 && richter < 10:
		magnitude = "great"
	default:
		magnitude = "massive"
	}
	fmt.Printf("%v is %v.\n", input, magnitude)

}
