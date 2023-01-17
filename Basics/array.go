package main

import "fmt"

// Global Array
var strArr []string

func arrayMaipulation() {
	fmt.Println(strArr)
	strArr = append(strArr, "Dibesh")
	//  Local Array
	intArr := []int{1, 2, 3}
	fmt.Println(intArr)
}

func main() {
	strArr = append(strArr, "Dinesh")
	arrayMaipulation()
	fmt.Println(strArr)
	for i, v := range strArr {
		println("\t", i, v)
	}
}
