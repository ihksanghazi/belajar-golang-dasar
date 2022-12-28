package main

import "fmt"

// Menentukan struct bernama Person
type Person struct {
	Name string
	Age  int
}

// Menambahkan method bernama Greet pada struct Person
func (p Person) Greet() {
	fmt.Println("Hello, my name is", p.Name)
}

func main() {
	// Membuat instance dari struct Person
	person := Person{Name: "John", Age: 30}

	// Memanggil method Greet
	person.Greet() // Output: Hello, my name is John
}
