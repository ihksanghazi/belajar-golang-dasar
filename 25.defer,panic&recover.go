package main

import "fmt"

func main() {
	// Menambahkan defer untuk mencetak pesan
	defer fmt.Println("Defer executed!")

	// Mencoba mengakses indeks yang tidak valid pada slice
	s := []int{1, 2, 3}
	fmt.Println(s[10]) // Akan menyebabkan panic

	// Menambahkan recover untuk menangani panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
}
