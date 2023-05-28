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
// EXERCISE: Break Up
//
//  1. Extend the "Only Evens" exercise
//  2. This time, use an infinite loop.
//  3. Break the loop when it reaches to the `max`.
//
// RESTRICTIONS
//  You should use the `break` statement once.
//
// HINT
//  Do not forget incrementing the `i` before the `continue`
//  statement and at the end of the loop.
//
// EXPECTED OUTPUT
//    go run main.go 1 10
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
	for {
		if min <= max {
			if min%2 == 0 {
				if min != max {
					output += fmt.Sprintf("%d + ", min)
				} else {
					sum += min
					output += fmt.Sprintf("%d = %d", min, sum)
				}
				sum += min
			}
		} else {
			break
		}
		min++
	}

	fmt.Println(output)
}
