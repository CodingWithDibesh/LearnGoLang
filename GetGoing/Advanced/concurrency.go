package main

import (
	"fmt"
	"time"
)

func heavy() {
	time.Sleep(time.Second * 5)
	fmt.Println("Stopped")
}

func main() {
	fmt.Println("GoRoutines Initialized")
	go heavy()
	fmt.Println("Completed")
	select {}
}
