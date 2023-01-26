package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func typeCasting() {
	strNumber := "1000"
	intNum, err := strconv.Atoi(strNumber)
	fmt.Println(intNum, err, reflect.TypeOf(intNum))
	// Format -> x to string
	// Parse -> string to x
}
