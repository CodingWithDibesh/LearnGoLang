package main

import (
	"fmt"
)

func main() {
	fmt.Println("Consuming Channels Buffered")
	c := make(chan int, 3)

	go func() {
		c <- 1
		c <- 2
		c <- 4
		c <- 5
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}

}
