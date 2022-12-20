package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan Ke - ", i)
	}

	//Dalam kode di atas, terdapat perulangan for yang menjalankan sebuah blok kode 10 kali. Setiap kali perulangan dijalankan, variabel i akan bertambah 1. Setelah perulangan ke-10, perulangan akan berhenti.
	//Di dalam blok kode perulangan, terdapat perintah if yang mengecek apakah nilai dari variabel i sama dengan 5. Jika benar, maka perintah break akan dijalankan, yang akan menghentikan perulangan yang sedang berlangsung.
	//Setelah itu, terdapat baris kode yang mencetak string "Perulangan Ke - " dan nilai dari variabel i ke layar. Baris kode ini akan dijalankan setiap kali perulangan dijalankan, sehingga akan tercetak 10 kali ke layar. Namun, karena perulangan akan berhenti setelah perulangan ke-5 karena perintah break, maka hanya akan tercetak 5 kali ke layar.

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Perulangan Ke - ", i)
	}

	//Perulangan ini akan menjalankan blok kode di dalamnya sebanyak 10 kali. Pada setiap iterasi, variabel i akan bertambah 1.
	//Pada setiap iterasi, ada sebuah kondisi if yang memeriksa apakah nilai dari i adalah genap atau tidak. Jika nilainya genap, maka perintah continue akan dijalankan, yang akan membuat program langsung melanjutkan ke iterasi berikutnya tanpa mengeksekusi apapun yang ada di bawahnya. Jika nilainya tidak genap, maka perintah fmt.Println akan dijalankan dan akan mencetak "Perulangan Ke - i" ke layar.

}
