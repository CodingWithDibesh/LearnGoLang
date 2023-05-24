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
// EXERCISE: Only Evens
//
//  1. Extend the "Sum up to N" exercise
//  2. Sum only the even numbers
//
// RESTRICTIONS
//  Skip odd numbers using the `continue` statement
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("Please enter min and max")
		fmt.Println("[min] [max]")
		return
	}

	min, errMin := strconv.Atoi(args[1])
	max, errMax := strconv.Atoi(args[2])

	if errMin != nil && errMax != nil {
		fmt.Println("[min] [max] must be numeric value")
		return
	}

	if min > max {
		fmt.Println("[min] value must be greater than [max]")
		return
	}

	output, sum := "", 0
	for ; min <= max; min++ {
		if min%2 == 0 {
			if min != max {
				output += fmt.Sprintf("%d + ", min)
			} else {
				sum += min
				output += fmt.Sprintf("%d = %d", min, sum)
			}
			sum += min
		}
	}

	fmt.Println(output)
}
