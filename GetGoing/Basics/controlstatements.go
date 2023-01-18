package main

import "fmt"

func main() {
	fmt.Println("Control Statements")
	// If/Else, For, Switch, Break, Continue
	f := true
	flag := &f
	if flag == nil {
		fmt.Println("Value is Nil")
	} else if *flag {
		fmt.Println("Value is true")
	} else {
		fmt.Println("Value is False")
	}

	// Looping Over Number
	for i := 0; i < 10; i++ {
		fmt.Printf("%d is in the loop\n", i)
	}

	myMap := make(map[string]interface{})
	myMap["Name"] = "Dibesh"
	myMap["Age"] = 12

	for k, v := range myMap {
		fmt.Printf("Key:%s and Value:%v\n", k, v)
	}

	for i := 0; i < 200; i++ {
		if i%2 == 0 {
			continue
		} else if i == 149 {
			fmt.Println("Loop has limit of", i)
			break
		}
		fmt.Println(i, "is a odd number")
	}

	getDay(2)
	getDay(7)
	getDay(9)

}

func getDay(day int) {
	switch day {
	case 1:
		fmt.Println("Sunday")
		break
	case 2:
		fmt.Println("Monday")
		break
	case 3:
		fmt.Println("Tuesday")
		break
	case 4:
		fmt.Println("Wednesday")
		break
	case 5:
		fmt.Println("ThursDay")
		break
	case 6:
		fmt.Println("Friday")
		break
	case 7:
		fmt.Println("Saturday")
		break
	default:
		fmt.Println("Invalid Number")
		break
	}
}
