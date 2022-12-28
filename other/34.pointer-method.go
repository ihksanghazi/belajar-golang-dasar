package main

import "fmt"

type Man struct {
	name string
}

func (a *Man) Married() {
	a.name = "Mr. " + a.name
}

func main() {
	sandy := Man{name: "Sandy"}
	sandy.Married()

	fmt.Println(sandy)
}
