package main

import "fmt"

func main() {
	var workOut = map[string]func(){
		"Variables":    variables,
		"Constants":    constants,
		"TypeCasting":  typeCasting,
		"Operators":    operators,
		"Conditionals": conditionals,
		"Loops":        loops,
	}
	fmt.Println(workOut)

	for index, value := range workOut {
		fmt.Println("\n\tWorking with ", index)
		value()
	}

}
