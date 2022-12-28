package main

import "fmt"

func main() {

	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan Ke = ", counter)
		counter++
	}

	// Pada baris pertama, variabel counter diberi nilai 1. Kemudian, pada baris kedua terdapat perulangan for yang akan mengevaluasi apakah nilai variabel counter kurang dari atau sama dengan 10. Jika kondisi tersebut terpenuhi, maka akan mencetak string "Perulangan Ke = " dan nilai dari variabel counter ke layar.
	// Setelah itu, pada baris ketiga, nilai dari variabel counter akan ditambah 1 menggunakan operator ++. Kemudian, perulangan akan kembali ke baris kedua dan mengevaluasi kondisi yang sama.

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("perulangan ke = ", counter)
	}

	// Pada baris pertama, terdapat perulangan for yang terdiri dari tiga bagian: inisialisasi, kondisi, dan increment. Pada bagian inisialisasi, variabel counter diberi nilai 1. Pada bagian kondisi, akan diuji apakah nilai variabel counter kurang dari atau sama dengan 10. Pada bagian increment, nilai dari variabel counter akan ditambah 1 menggunakan operator ++.
	// Jika kondisi tersebut terpenuhi, maka akan mencetak string "perulangan ke = " dan nilai dari variabel counter ke layar. Kemudian, perulangan akan kembali ke bagian increment dan menambah nilai dari variabel counter.

	slice := []string{"Nur", "Sandy", "Ihksan"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// Pada bagian pertama, variabel slice dideklarasikan sebagai sebuah slice (array dinamis) bertipe string yang berisi tiga elemen string. Elemen-elemen tersebut adalah "Nur", "Sandy", dan "Ihksan".
	// Pada bagian kedua, terdapat perulangan for yang akan mencetak elemen-elemen dari slice tersebut ke layar. Perulangan tersebut akan dijalankan sebanyak panjang slice (len(slice)), yang dihitung dengan menghitung jumlah elemen dari slice tersebut. Setiap iterasi perulangan, elemen dari slice akan diakses dengan menggunakan indeks yang sesuai (i) dan dicetak ke layar menggunakan perintah fmt.Println().

	names := []string{"Azhi", "Muahnaf", "Rais"}

	for indek, name := range names {
		fmt.Println("indek = ", indek, "name = ", name)
	}

	// Pada bagian pertama, variabel names dideklarasikan sebagai sebuah slice (array dinamis) bertipe string yang berisi tiga elemen string. Elemen-elemen tersebut adalah "Azhi", "Muahnaf", dan "Rais".
	// Pada bagian kedua, terdapat perulangan for yang akan mengakses elemen-elemen dari slice tersebut dan mencetak indeks dan nilai elemen tersebut ke layar. Perulangan tersebut menggunakan konstruksi range names, yang akan mengembalikan indeks dan nilai elemen untuk setiap iterasi perulangan. Pada setiap iterasi, indeks (indek) dan nilai (name) elemen akan dicetak ke layar menggunakan perintah fmt.Println().
}
