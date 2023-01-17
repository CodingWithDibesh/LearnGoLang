package main

import "fmt"

// Data Encapsulations as similar to C++
type Human struct {
	name string
	age  int
}

func (h Human) Print() {
	fmt.Println(h.name)
}

func (h Human) Introduce() {
	fmt.Println("Hi I'm", h.name)
	fmt.Println("I'm", h.age, "years old!")
}

func (h Human) GetName() string {
	return h.name
}

func main() {
	fmt.Println("Testing Structures")
	dibesh := Human{"Dibesh", 12}
	fmt.Println(dibesh)
	dibesh.Print()
	dibesh.Introduce()
	fmt.Println(dibesh.GetName())
}
