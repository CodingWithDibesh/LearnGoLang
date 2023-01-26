package main

import (
	"fmt"
	"reflect"
)

// Constant with defined type
const LAPTOP string = "HP-Omen"

// Constant with implicit type definition
const PRICE = 124000

func constants() {
	// Syntax: const identifier type(optional) = value
	// Grouped Constant Declaration
	const (
		PRODUCT = "Mackintosh"
		InStock = false
		InNRS   = 12.34
	)
	fmt.Println(LAPTOP, PRICE)
	fmt.Println(PRODUCT, InStock)
	fmt.Println("LAPTOP is of type", reflect.TypeOf(LAPTOP), "and a constant")
	fmt.Println("InNRS is of type", reflect.TypeOf(InNRS), "and a constant")
	fmt.Println("PRICE is of type", reflect.TypeOf(PRICE), "and a constant")
	fmt.Println("InStock is of type", reflect.TypeOf(InStock), "and a constant")

}
