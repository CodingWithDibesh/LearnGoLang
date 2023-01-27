package main

import (
	"fmt"
)

func loops() {
	iterations := 2
	characterSeq := "NEPAL"
	fmt.Println("Enter the number of iterations:", iterations)
	// fmt.Scan(&iterations)
	fmt.Println("Enter the character sequence:", characterSeq)
	// fmt.Scan(&characterSeq)
	fmt.Println("Printing sequence: ")
	length := len(characterSeq)
	for i := 0; i < iterations; i++ {
		for j := 0; j < length; j++ {
			for k := 0; k <= j; k++ {
				//  The Value here is of type rune and needs to be parsed as string
				value := characterSeq[k]
				fmt.Print(" ", string(value))
			}
			fmt.Println()
		}
	}
	fmt.Println("\n\tPlaying with rune and strings")
	for index, value := range characterSeq {
		fmt.Print("Index:", index, ", has value ", value, ", which is :", string(value), "\n")
	}
}
