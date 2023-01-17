package main

import "fmt"

type Human interface {
	Walk()
	Talk()
}

type Man struct {
	name string
}

type Woman struct {
	name string
}

func NewMan(arg string) Human {
	return &Man{arg}
}

func (m *Man) Walk() {
	fmt.Println("He is waling")
	fmt.Println(m.name, "is walking.")
}

func (m *Man) Talk() {
	fmt.Println(m.name, "is talking.")
}

func (w *Woman) Walk() {
	fmt.Println("She is waling")
	fmt.Println(w.name, "is walking.")
}

func main() {
	fmt.Println("Playing with Interfaces")
	dina := Woman{"dina"}
	john := NewMan("John")
	dina.Walk()
	john.Walk()
	john.Talk()
}
