package main

import "fmt"

func main() {
	// Menggunakan anonymous function sebagai variabel
	add := func(x, y int) int {
		return x + y
	}

	result := add(1, 2)
	fmt.Println(result) // Output: 3

	// Menggunakan anonymous function sebagai parameter
	func() {
		fmt.Println("Hello, World!")
	}()
}
