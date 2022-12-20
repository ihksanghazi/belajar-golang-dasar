package main

import "fmt"

func main() {
	// Slice
	// * Tipe Data Slice adalah Potongan Data Dari Array
	// * Slice Mirip Array yang membedakan adalah ukuran Slice Bisa Berubah
	// * Slice dan Array selalu terkoneksi, dimana slice adalah data yang mengakses sebagian atau seluruh data di array

	// Detail Tipe Data Slice
	// * Tipe Data Slice memiliki 3 data, yaitu pointer, length, dan capaciti
	// * pointer adalah penunjuk data pertama di array para slice
	// * length adalah panjang dari slice
	// * Capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity

	// Membuat Slice dari array

	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	//Ini adalah sebuah contoh deklarasi array di bahasa pemrograman Go yang menggunakan sintaks shorthand dengan tanda titik-tiga (...). Array months terdiri dari elemen-elemen bertipe string yang menyimpan nama-nama bulan dalam bahasa Indonesia.
	//Tanda titik-tiga (...) digunakan untuk menandakan bahwa panjang dari array akan dihitung otomatis sesuai dengan jumlah elemen yang dideklarasikan. Dengan kata lain, panjang dari array months akan secara otomatis ditentukan menjadi 12 karena terdapat 12 elemen yang dideklarasikan.

	var slice1 = months[4:7]

	//Untuk membuat slice dari sebuah array, gunakan notasi [awal:akhir] setelah nama array. Indeks awal dan akhir merupakan indeks dari elemen yang akan disimpan dalam slice. Indeks awal dari slice tersebut adalah 4, sedangkan indeks akhir adalah 7. Sehingga, slice slice1 akan menyimpan elemen ke-5 sampai ke-7 dari array months, yaitu "Mei", "Juni", dan "Juli".

	fmt.Println("Isi slice1 = ", slice1)
	fmt.Println("Panjang slice1 = ", len(slice1))
	fmt.Println("Capacity slice1 = ", cap(slice1))

	// Jika Array diubah maka slice juga akan berubah
	// months[5] = "diubah"
	// fmt.Println(slice1)

	// Jika slice diubah maka array juga akan berubah
	// slice1[0] = "Mei Lagi"
	// fmt.Println(months)

	// Membuat slice2 dari index 10 sampai terakhir

	var slice2 = months[10:]

	//Indeks awal dari slice tersebut adalah 10, sedangkan indeks akhir tidak ditentukan. Sehingga, slice slice2 akan menyimpan elemen ke-11 dan ke-12 dari array months, yaitu "November" dan "Desember".

	fmt.Println("isi slice2 = ", slice2)

	// membuat slice3 dengan cara append dari slice2
	var slice3 = append(slice2, "Sandy")

	//Dalam kode di atas, terdapat variabel slice2 yang merupakan sebuah slice bertipe string. Kemudian, terdapat deklarasi variabel slice3 yang merupakan hasil dari pemanggilan fungsi append() yang menerima 2 argument, yaitu slice2 dan string "Sandy". Fungsi append() akan menambah elemen "Sandy" ke dalam slice slice2 dan mengembalikan slice yang baru yang berisi elemen dari slice2 dan elemen "Sandy".

	fmt.Println("isi slice3 = ", slice3)

	// mengubah isi slice3
	slice3[1] = "Bukan December"

	// isi slice3 berubah
	fmt.Println("isi slice3 = ", slice3)
	// isi slice2 tidak berubah
	fmt.Println("isi slice2 = ", slice2)
	// isi array month tidak berubah
	fmt.Println("isi array month = ", months)

	//append akan membuat slice baru dengan menambah data ke posisi terakhir slice, jika kapasitas sudah penuh, maka akan membuat array baru

	// Membuat Slice dengan method make
	newslice := make([]string, 2, 5)
	newslice[0] = "Nur"
	newslice[1] = "Sandy"

	//Dalam kode di atas, terdapat deklarasi variabel newslice yang merupakan hasil pemanggilan fungsi make() dengan 3 argument. Argumen pertama merupakan tipe data dari elemen slice, yaitu string. Argumen kedua merupakan panjang dari slice yang akan dibuat, yaitu 2. Argumen ketiga merupakan kapasitas dari slice yang akan dibuat, yaitu 5.

	fmt.Println("Membuat slice dengan method make = ", newslice)
	fmt.Println("panjang Dari newSlice = ", len(newslice))
	fmt.Println("capacity dari newslice = ", cap(newslice))

	copyslice := make([]string, len(newslice), cap(newslice))
	copy(copyslice, newslice)
	fmt.Println("Membuat copyslice dengan method copy", copyslice)

	//Dalam kode di atas, terdapat variabel newslice yang merupakan sebuah slice bertipe string. Kemudian, terdapat deklarasi variabel copyslice yang merupakan hasil pemanggilan fungsi make() dengan 3 argument. Argumen pertama merupakan tipe data dari elemen slice, yaitu string. Argumen kedua merupakan panjang dari slice yang akan dibuat, yaitu panjang dari slice newslice. Argumen ketiga merupakan kapasitas dari slice yang akan dibuat, yaitu kapasitas dari slice newslice.
	//Kemudian, terdapat pemanggilan fungsi copy() yang menerima 2 argument, yaitu copyslice dan newslice. Fungsi copy() akan menyalin elemen dari slice newslice ke slice copyslice dan mengembalikan jumlah elemen yang disalin. Setelah itu, slice copyslice akan terdiri dari 2 elemen bertipe string yang sama dengan elemen dari slice newslice, yaitu "Nur" dan "Sandy".

	// Hati Hati saat membuat array dan slice
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
}
