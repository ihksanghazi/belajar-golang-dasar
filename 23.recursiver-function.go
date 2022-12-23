package main

import "fmt"

// Function untuk mencari faktorial dari sebuah angka
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	result := factorial(5)
	fmt.Println(result) // Output: 120
}
