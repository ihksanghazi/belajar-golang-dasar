package main

import "fmt"

func main() {
	var name = "Sandy"

	if name == "Sand" {
		fmt.Println("hello udin")
	} else if name == "udin" {
		fmt.Print("hello udin")
	} else {
		fmt.Println("hello")
	}

	// 	Pada baris pertama, variabel name diberi nilai "Sandy". Kemudian, pada baris kedua terdapat percabangan if, yang akan mengevaluasi apakah nilai variabel name sama dengan string "Sand". Jika kondisi tersebut terpenuhi, maka akan mencetak string "hello udin" ke layar.
	// Jika kondisi tersebut tidak terpenuhi, maka akan diikuti oleh percabangan else if pada baris ketiga, yang akan mengevaluasi apakah nilai variabel name sama dengan string "udin". Jika kondisi tersebut terpenuhi, maka akan mencetak string "hello udin" ke layar.
	// Jika kedua kondisi tersebut tidak terpenuhi, maka akan diikuti oleh percabangan else pada baris keempat, yang akan mencetak string "hello" ke layar.

	if length := len(name); length > 5 {
		fmt.Println("Nama Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}

	// 	Pada baris pertama, terdapat percabangan if yang mengevaluasi panjang string dari variabel name. Panjang string tersebut dihitung menggunakan fungsi len(), dan hasilnya disimpan pada variabel length. Kemudian, variabel length akan diuji apakah lebih besar dari 5. Jika kondisi tersebut terpenuhi, maka akan mencetak string "Nama Terlalu Panjang" ke layar.

	// Jika kondisi tersebut tidak terpenuhi, maka akan diikuti oleh percabangan else pada baris kedua, yang akan mencetak string "Nama Sudah Benar" ke layar.
}
