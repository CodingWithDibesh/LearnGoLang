package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Consuming Channels")
	c := make(chan int)

	// Send

	go func() {
		c <- 1
	}()

	// Sniff
	val := <-c
	fmt.Println(val)
	go func() {
		c <- 5
	}()

	time.Sleep(time.Second * 2)
	val = <-c
	fmt.Println(val)

}
