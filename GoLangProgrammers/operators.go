package main

import "fmt"

func operators() {
	var x, y = 35, 7

	fmt.Printf("x + y = %d\n", x+y)
	fmt.Printf("x - y = %d\n", x-y)
	fmt.Printf("x * y = %d\n", x*y)
	fmt.Printf("x / y = %d\n", x/y)
	fmt.Printf("x mod y = %d\n", x%y)

	x, y = 15, 25

	x++
	fmt.Printf("x++ = %d\n", x)

	y--
	fmt.Printf("y-- = %d\n", y)

	x = y
	fmt.Println("= ", x)

	x = 15
	x += y
	fmt.Println("+=", x)

	x = 50
	x -= y
	fmt.Println("-=", x)

	x = 2
	x *= y
	fmt.Println("*=", x)

	x = 100
	x /= y
	fmt.Println("/=", x)

	x = 40
	x %= y
	fmt.Println("%=", x)

	x, y = 15, 25

	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x < y)
	fmt.Println(x <= y)
	fmt.Println(x > y)
	fmt.Println(x >= y)

	x, y = 10, 20
	z := 30

	fmt.Println(x < y && x > z)
	fmt.Println(x < y || x > z)
	fmt.Println(!(x == y && x > z))

	var X uint = 9  //0000 1001
	var Y uint = 65 //0100 0001
	var Z uint

	Z = X & Y
	fmt.Println("X & Y  =", Z)

	Z = X | Y
	fmt.Println("X | Y  =", Z)

	Z = X ^ Y
	fmt.Println("X ^ Y  =", Z)

	Z = X << 1
	fmt.Println("X << 1 =", Z)

	Z = X >> 1
	fmt.Println("X >> 1 =", Z)
}
