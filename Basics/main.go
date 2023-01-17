// Defining the package name
package main

// Importing I/O package
import "fmt"

// Executing Main Function
func main() {
	// Using Println to print the message in console.
	fmt.Println("Hello World")
	// Defining and assigning a value to a variable
	var val1 = 2
	fmt.Println(val1)

	// Defining and assigning multiple numeric variables
	var (
		val2 int32 = 3
		val3 int64 = 4
	)

	// TypeCasting and adding up the variables
	fmt.Println(int64(val2) + val3)

	// Shorthand Variable Declaration
	val4 := 5
	fmt.Printf("val4 has value %d and is of type %T\n", val4, val4)

	// Declaring a string via shorthand
	val5 := "MyName is Dibesh"
	fmt.Println(val5)
	// Checking string Mutability
	val5 = "My Name is Dibesh"
	fmt.Println(val5)

}
