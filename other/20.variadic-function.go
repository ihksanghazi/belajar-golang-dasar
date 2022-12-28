package main

import "fmt"

func add(numbers ...int) int {
	// Inisialisasi variabel untuk menyimpan hasil penjumlahan
	sum := 0
	// Perulangan untuk menambahkan setiap angka
	for _, number := range numbers {
		sum += number
	}
	// Mengembalikan hasil penjumlahan
	return sum
}

func main() {
	result1 := add(1, 2, 3, 4, 5)
	fmt.Println(result1) // Output: 15

	// Menggunakan variadic function dengan slice
	numbers := []int{6, 7, 8, 9, 10}
	result2 := add(numbers...)
	fmt.Println(result2) // Output: 40
}
