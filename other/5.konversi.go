package main

import "fmt"

func main() {

	//Type casting adalah proses mengubah tipe data suatu nilai ke tipe data lain.
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	//Pada baris pertama, kita mendeklarasikan variabel nilai32 dengan tipe data int32 dan memberinya nilai 100000. Tipe data int32 adalah tipe data integer dengan panjang 32 bit.
	// Pada baris kedua, kita mendeklarasikan variabel nilai64 dengan tipe data int64 dan mengisi nilainya dengan hasil type casting dari variabel nilai32. Kita menggunakan tipe data int64 untuk menampung nilai yang lebih besar dari batas maksimum tipe data int32.
	// Pada baris ketiga, kita mendeklarasikan variabel nilai8 dengan tipe data int8 dan mengisi nilainya dengan hasil type casting dari variabel nilai32. Tipe data int8 hanya dapat menampung nilai integer dengan rentang -128 sampai 127, sehingga nilai 100000 tidak dapat ditampung oleh tipe data ini.

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "Nursandy Ihksan"
	var e = name[0]
	var eString = string(e)

	//Pada baris pertama, kita mendeklarasikan variabel name dengan tipe data string dan memberinya nilai "Nursandy Ihksan".
	//Pada baris kedua, kita mendeklarasikan variabel e dan mengisi nilainya dengan mengakses elemen pertama dari variabel name menggunakan indexing. Indexing merupakan cara untuk mengakses elemen-elemen dari tipe data yang memiliki elemen seperti array, slice, dan string. Indexing dimulai dari 0, sehingga elemen pertama dari variabel name adalah "N".
	//Pada baris ketiga, kita mendeklarasikan variabel eString dengan tipe data string dan mengisi nilainya dengan hasil type casting dari variabel e. Variabel e adalah tipe data byte, sedangkan tipe data string hanya dapat menampung nilai-nilai string, sehingga kita perlu melakukan type casting agar nilai e dapat ditampung oleh tipe data string.

	fmt.Println(name)
	fmt.Println(eString)

}
