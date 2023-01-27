package main

import (
	"fmt"
	"math"
	"os"
)

func functions() {
	printName("Dibesh")
	fmt.Println("2 + 222 =", add(2, 222))
	fmt.Println("Value of PI is", pi())
	whoAmI()
	firstName, middleName, lastName := introduceYourSelf()
	fmt.Println("My name is", firstName, middleName, lastName)
	name := "Dibesh"
	age := 10
	fmt.Printf("Value Before => Name: %s and Age: %d\n", name, age)
	pointersAreGreat(&name, &age)
	fmt.Printf("Value After => Name: %s and Age: %d\n", name, age)
	fmt.Printf("Area of my house is : %d meeter square.\n", area(20, 30))
	// GoLang also supports IIFE -> Immediately Invoked Function Expression
	func(name string) {
		fmt.Println("Well It's IIFE printing Name:", name)
	}("Dibesh")
	// Closer in use as well
	fmt.Printf(
		"%s says 100 (°F) = %.2f (°C)\n", name,
		func(f float64) float64 {
			return (f - 32.0) * (5.0 / 9.0)
		}(100),
	)
	// Anonymous Function accessing the variable inside loop
	for i := 10.0; i < 100; i += 10.0 {
		rad := func() float64 {
			return i * 39.370
		}()
		fmt.Printf("%.2f Meter = %.2f Inch\n", i, rad)
	}
	fmt.Println("Executing my Super Function:")
	addFn := func() int64 {
		return add(10, 20)
	}
	fmt.Println("Result of my super function:", compounder(addFn)(3))
	num := squareSum(5)(6)(7)
	fmt.Println("Result:", num)
}

// Function with Argument and no return value
func printName(name string) {
	fmt.Println("Your Name is:", name)
}

// Function with Argument and return value
// Function with multiple params
func add(number1, number2 int64) int64 {
	return number1 + number2
}

// Function with no Argument and return value
func pi() (PI float64) {
	PI = math.Pi
	// Returning Statement without specifying var name
	return
}

// Function with no Argument and no return value
func whoAmI() {
	os, _ := os.Hostname()
	fmt.Println(os)
}

// Function returning multiple values
func introduceYourSelf() (firstName, middleName, lastName string) {
	firstName = "Dibesh"
	middleName = "Raj"
	lastName = "Subedi"
	return
}

func pointersAreGreat(name *string, age *int) {
	*name += " Raj Subedi" // defrencing pointer address
	*age += 15
}

// Anonymous Function
var area = func(length, breadth int) int {
	return length * breadth
}

// High Order Function
// Note: A Function that receives a function as an argument or returns the function as output.
// Func returning a Func
func compounder(magicFunc func() int64) func(product int64) float64 {
	return func(product int64) float64 {
		value := magicFunc()
		return float64(value) * float64(product)
	}
}

//  Defining Custom Function Type

type Number func() float64
type InputNumber func(string) float64

type First func(int) int
type Second func(int) First

func squareSum(x int) Second {
	return func(y int) First {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}
