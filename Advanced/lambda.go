package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Lambda/ Anonymous functions")
	//  Wait Group for GoRoutines
	wg := &sync.WaitGroup{}

	// mg:=&sync.Mutex{}
	// Use Mutex when using shared resources
	// mg.Lock()
	// mg.Unlock()
	// mg.TryLock()

	wg.Add(1)

	go func() {
		fmt.Println("I'm a Lambda Function")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Its an end")
}
