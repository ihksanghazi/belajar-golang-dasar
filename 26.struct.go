package main

import "fmt"

// Menentukan struct bernama Person
type Person struct {
	Name string
	Age  int
}

func main() {
	// Membuat instance dari struct Person
	person1 := Person{Name: "John", Age: 30}
	person2 := Person{"Alice", 25}

	// Mencetak nilai dari struct
	fmt.Println(person1.Name) // Output: John
	fmt.Println(person2.Age)  // Output: 25
}
