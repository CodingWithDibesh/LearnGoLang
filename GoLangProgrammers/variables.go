package main

import (
	"fmt"
	"reflect"
)

func variables() {
	//  Syntax: var identifier type(optional) = value
	// Variable Declaration With Initialization
	var bio string = "I'm Dibesh Raj Subedi"
	// Multiple Variable Declaration With initialization
	var firstName, middleName, LastName string = "Dibesh", "Raj", "Subedi"
	// Short Variable Declaration with Implicitly Type Definition
	age := 24
	isWorking := true
	// Grouped Variable Declaration
	var (
		location string = "Gaushala"
		zipCode  int    = 44660
	)
	// Printing the values in console
	fmt.Println(bio)
	fmt.Println(firstName, middleName, LastName, age, isWorking)
	fmt.Println(location, zipCode)

	// Checking Type of Variables
	fmt.Println("firstName is of type ", reflect.TypeOf(firstName))
	fmt.Println("age is of type ", reflect.TypeOf(age))
	fmt.Println("isWorking is of type ", reflect.TypeOf(isWorking))
}
