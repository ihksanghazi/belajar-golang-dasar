package main

import "fmt"

func main() {
	var name1 = "Nur"
	var name2 = "Sandy"

	var result = name1 == name2

	//Pada baris pertama dan kedua, kita mendeklarasikan variabel name1 dengan nilai "Nur" dan variabel name2 dengan nilai "Sandy".
	//Pada baris keempat, kita mendeklarasikan variabel result dan mengisi nilainya dengan hasil perbandingan dari variabel name1 dan name2 menggunakan operator perbandingan sama dengan. Sehingga nilai variabel result adalah false karena nilai dari kedua variabel tersebut tidak sama.

	fmt.Println(result)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)

	//Code di atas merupakan contoh penggunaan operator perbandingan lebih besar dari (>), lebih kecil dari (<), sama dengan (==), dan tidak sama dengan (!=)
}
