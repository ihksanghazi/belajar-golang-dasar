package main

import "fmt"

func main() {
	var result = 10 + 10

	//kita mendeklarasikan variabel result dan mengisi nilainya dengan hasil penjumlahan dari 10 dan 10 menggunakan operator penjumlahan.

	fmt.Println(result)

	var a = 10
	var b = 20
	var c = a + b
	fmt.Println(c)

	//Pada baris pertama dan kedua, kita mendeklarasikan variabel a dengan nilai 10 dan variabel b dengan nilai 20.
	//Pada baris ketiga, kita mendeklarasikan variabel c dan mengisi nilainya dengan hasil penjumlahan dari variabel a dan b menggunakan operator penjumlahan.

	var i = 10
	i += 10

	//Pada baris pertama, kita mendeklarasikan variabel i dengan nilai 10.
	//Pada baris kedua, kita menggunakan operator penambahan sama dengan untuk menambahkan 10 ke variabel i dan menyimpannya kembali ke variabel i. Sehingga nilai variabel i akan menjadi 20.

	fmt.Println(i)

	i++
	fmt.Println(i)

	//Pada baris pertama, kita menggunakan operator increment untuk menambahkan 1 ke variabel i. Sehingga nilai variabel i akan menjadi 21.

	var negatif = -100
	var positif = +100

	//Pada baris pertama, kita mendeklarasikan variabel negatif dan memberinya nilai -100 menggunakan operator unary minus. Pada baris kedua, kita mendeklarasikan variabel positif dan memberinya nilai 100 menggunakan operator unary plus.

	fmt.Println(negatif)
	fmt.Println(positif)
}
