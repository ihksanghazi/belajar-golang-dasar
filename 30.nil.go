package main

import "fmt"

func main() {
	// Membuat variabel bertipe pointer dengan nilai awal nil
	var p *int
	fmt.Println(p) // mencetak nil

	// Menugaskan nilai ke variabel p
	n := 10
	p = &n
	fmt.Println(p) // mencetak alamat memori dari variabel n

	// Menugaskan nil ke variabel p
	p = nil
	fmt.Println(p) // mencetak nil
}
