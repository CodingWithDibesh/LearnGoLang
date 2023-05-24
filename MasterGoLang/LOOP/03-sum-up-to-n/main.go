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
// EXERCISE: Sum up to N
//
//  1. Get two numbers from the command-line: min and max
//  2. Convert them to integers (using Atoi)
//  3. By using a loop, sum the numbers between min and max
//
// RESTRICTIONS
//  1. Be sure to handle the errors. So, if a user doesn't
//     pass enough arguments or she passes non-numerics,
//     then warn the user and exit from the program.
//
//  2. Also, check that, min < max.
//
// HINT
//  For converting the numbers, you can use `strconv.Atoi`.
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 = 55
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
		if min != max {
			output += fmt.Sprintf("%d + ", min)
		} else {
			sum += min
			output += fmt.Sprintf("%d = %d", min, sum)
		}
		sum += min
	}

	fmt.Println(output)
}
