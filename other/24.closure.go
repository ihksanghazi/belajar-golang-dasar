package main

import "fmt"

func adder() func(int) int {
	// Inisialisasi variabel sum dengan nilai 0
	sum := 0
	// Mengembalikan function yang menambahkan nilai ke sum
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// Memanggil closure
	add := adder()
	// Menambahkan nilai ke sum
	fmt.Println(add(1)) // Output: 1
	fmt.Println(add(2)) // Output: 3
	fmt.Println(add(3)) // Output: 6
}
