package main

import "fmt"

func addSubtract(x, y int) (add int, subtract int) {
	// Menambahkan dua angka
	add = x + y
	// Mengurangi dua angka
	subtract = x - y
	// Mengembalikan hasil dari penambahan dan pengurangan
	return
}

func main() {
	a, s := addSubtract(5, 3)
	fmt.Println("Hasil penambahan:", a)  // Output: Hasil penambahan: 8
	fmt.Println("Hasil pengurangan:", s) // Output: Hasil pengurangan: 2
}
