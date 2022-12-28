package main

import "fmt"

type hasName interface {
	getname() string
}

func sayHello(h hasName) {
	fmt.Println("hello", h.getname())
}

type person struct {
	name string
}

func (p person) getname() string {
	return p.name
}

type animal struct {
	name string
}

func (an animal) getname() string {
	return an.name
}

func main() {
	var sandy person
	sandy.name = "Sandy"
	sayHello(sandy)

	var binatang animal
	binatang.name = "Kucing"
	sayHello(binatang)
}
