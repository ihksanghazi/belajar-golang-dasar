package main

import "fmt"

// Menentukan interface kosong bernama Any
type Any interface{}

func main() {
	// Menyimpan tipe data apapun di dalam interface kosong
	var a Any = 10
	fmt.Println(a) // Output: 10

	a = "hello"
	fmt.Println(a) // Output: hello

	a = true
	fmt.Println(a) // Output: true
}
