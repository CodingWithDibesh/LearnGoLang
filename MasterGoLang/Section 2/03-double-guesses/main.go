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
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Double Guesses
//
//  Let the player guess two numbers instead of one.
//
// HINT:
//  Generate random numbers using the greatest number
//  among the guessed numbers.
//
// EXAMPLES
//  go run main.go 5 6
//  Player wins if the random number is either 5 or 6.
// ---------------------------------------------------------

const (
	maxTurns = 5 // less is more difficult
	usage    = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers twice.
Your mission is to guess two numbers which will match those numbers.

The greater your number is, harder it gets.

Wanna play?
`
)

func main() {
	var guess []int

	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Printf(usage, maxTurns)
		return
	}

	for i := 0; i < 2; i++ {
		g, err := strconv.Atoi(args[i])
		if err != nil {
			fmt.Println("Not a number:", err)
			return
		}
		if g <= 0 {
			fmt.Printf("%d is not a positive number.\n", g)
			return
		}
		guess = append(guess, g)
	}

	for i := 1; i <= maxTurns; i++ {
		num1 := rand.Intn(guess[0]) + 1
		num2 := rand.Intn(guess[1]) + 1

		if num1 != guess[0] || num2 != guess[1] {
			continue
		}

		if num1 == guess[0] || num2 == guess[0] {
			switch rand.Intn(3) {
			case 0:
				fmt.Println("🎉  YOU WIN!")
			case 1:
				fmt.Println("🎉  YOU'RE AWESOME!")
			case 2:
				fmt.Println("🎉  PERFECT!")
			}
			return
		}
	}

	msg := "👻 YOU LOST..."
	switch rand.Intn(2) {
	case 0:
		msg += "Try again?"
	case 1:
		msg += "Better Luck Next Time 😜"
	}

	fmt.Println(msg)
}
