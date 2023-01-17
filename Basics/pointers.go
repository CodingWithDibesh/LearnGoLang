package main

import "fmt"

func main() {
	value1 := 20
	value2 := 5
	fmt.Println("Comparing Values of value1 and Value 2:", value1 == value2)
	fmt.Println("Comparing Address of value1 and Value 2:", &value1 == &value2)
	fmt.Println("Addresses:", &value1, &value2)
	fmt.Println(value1, value2)
	swap(&value1, &value2)
	fmt.Println(value1, value2)
}

func swap(value1, value2 *int) {
	var tmp int
	tmp = *value2
	*value2 = *value1
	*value1 = tmp
}
