package main

import "fmt"

// return multiple function
func getname() (string, string, string) {
	return "Nur", "Sandy", "Ihksan"
}

// named return values
func getCompleteName() (fistname, middlename, lastname string) {
	fistname = "Nur"
	middlename = "Sandy"
	lastname = "Ihksan"
	return fistname, middlename, lastname
}

// variadic function
// parameter tidak terbatas
func tambah(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {

	// multiple function
	Awal, Tengah, Akhir := getname()
	fmt.Println(Awal, Tengah, Akhir)

	_, sandy, ihksan := getname() // _ untuk menghiraukan return value pertama
	fmt.Println(sandy, ihksan)

	// named return values
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)

	// variadic function
	fmt.Println(tambah(1, 2, 3, 4, 5, 6, 7, 8, 9))

	angka := []int{12, 56, 8, 19, 26}
	fmt.Println(tambah(angka...))

}
