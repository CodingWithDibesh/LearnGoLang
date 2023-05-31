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
// EXERCISE: Verbose Mode
//
//  When the player runs the game like this:
//
//    go run main.go -v 4
//
//  Display each generated random number:

//    1 3 4 🎉  YOU WIN!
//
//  In this example, computer picks 1, 3, and 4. And the
//  player wins.
//
// HINT
//  You need to get and interpret the command-line arguments.
// ---------------------------------------------------------

const (
	maxTurns = 5 // less is more difficult
	usage    = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.
Your mission is to guess one of those numbers.
	
The greater your number is, harder it gets.
	
Wanna play?
	
(Provide -v flag to see the picked numbers.)
`
)

func main() {
	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Printf(usage, maxTurns)
		return
	}

	v := args[0] == "-v"
	guess, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Not a number:", err)
		return
	}

	if guess <= 0 {
		fmt.Printf("%d is not a positive number.\n", guess)
		return
	}

	ite := ""

	for i := 1; i <= maxTurns; i++ {
		num := rand.Intn(guess) + 1
		ite += " " + strconv.Itoa(num)
		if num != guess {
			continue
		}
		if num == guess {
			switch rand.Intn(3) {
			case 0:
				if v {
					fmt.Println(ite, "🎉  YOU WIN!")
				} else {
					fmt.Println("🎉  YOU WIN!")
				}
			case 1:
				if v {
					fmt.Println(ite, "🎉  YOU'RE AWESOME!")
				} else {
					fmt.Println("🎉  YOU'RE AWESOME!")
				}
			case 2:
				if v {
					fmt.Println(ite, "🎉  PERFECT!")
				} else {
					fmt.Println("🎉  PERFECT!")
				}
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

	if v {
		fmt.Println(ite, msg)
	} else {
		fmt.Println(msg)
	}
}
