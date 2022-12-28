package main

import "fmt"

func main() {

	// number
	fmt.Println("Satu", 1)
	fmt.Println("Dua", 2)
	fmt.Println("Tiga koma lima", 3.5)

	//Pada perintah pertama, string "Satu" dan nilai integer 1 akan dicetak ke layar. Pada perintah kedua, string "Dua" dan nilai integer 2 akan dicetak ke layar. Pada perintah ketiga, string "Tiga koma lima" dan nilai float 3.5 akan dicetak ke layar.

	// boolean
	fmt.Println("Benar = ", true)
	fmt.Println("Salah = ", false)

	//Pada perintah pertama, string "Benar = " dan nilai boolean true akan dicetak ke layar. Pada perintah kedua, string "Salah = " dan nilai boolean false akan dicetak ke layar.
	//Nilai boolean hanya memiliki dua kemungkinan, yaitu true (benar) atau false (salah). Nilai boolean sering digunakan dalam kondisi percabangan atau perulangan untuk menentukan apakah suatu statement harus dijalankan atau tidak.

	// string
	fmt.Println("Nur")
	fmt.Println("Nursandy")
	fmt.Println("Nursandy Ihksan")

	//Pada setiap perintah, hanya sebuah string yang akan dicetak ke layar. Pada perintah pertama, string "Nur" akan dicetak ke layar. Pada perintah kedua, string "Nursandy" akan dicetak ke layar. Pada perintah ketiga, string "Nursandy Ihksan" akan dicetak ke layar.

	fmt.Println(len("Nursandy Ihksan"))
	fmt.Println("Nursandy Ihksan"[0])

	//Pada perintah pertama, ekspresi len("Nursandy Ihksan") akan dihitung dan hasilnya akan dicetak ke layar. Fungsi len() akan mengembalikan panjang string yang diberikan sebagai argumen. Jadi, hasil yang akan dicetak adalah panjang string "Nursandy Ihksan", yaitu 15.
	//Pada perintah kedua, ekspresi "Nursandy Ihksan"[0] akan dihitung dan hasilnya akan dicetak ke layar. Dalam bahasa pemrograman Go, string merupakan sekumpulan dari karakter-karakter yang diurutkan secara berurutan, dan setiap karakter dapat diakses dengan menggunakan indeks. Indeks dimulai dari 0 untuk karakter pertama, 1 untuk karakter kedua, dan seterusnya. Jadi, ekspresi tersebut akan mengakses karakter pertama dari string "Nursandy Ihksan", yaitu karakter 'N'.
}
