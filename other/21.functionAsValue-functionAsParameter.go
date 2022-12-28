package main

import "fmt"

// Function pertama
func greet(name string) {
	fmt.Println("Hello, " + name)
}

// Function kedua
func add(x, y int) int {
	return x + y
}

func main() {
	// Menggunakan function sebagai variabel
	var f func(string)
	f = greet
	f("John") // Output: Hello, John

	// Menggunakan function sebagai parameter
	func(f func(int, int) int, x, y int) {
		result := f(x, y)
		fmt.Println(result)
	}(add, 1, 2) // Output: 3
}
