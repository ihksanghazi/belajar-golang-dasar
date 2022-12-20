package main

import "fmt"

func sayHello() {
	fmt.Println("Hello World")
}

//Fungsi sayHello adalah fungsi yang didefinisikan oleh pengguna dan tidak memiliki parameter masukan atau nilai kembali. Fungsi ini hanya mencetak "Hello World" ke layar dengan menggunakan perintah fmt.Println.

func sayHelloTo(nama string) {
	fmt.Println("Hello ", nama)
}

//Fungsi ini bernama sayHelloTo dan memiliki satu parameter bernama nama yang bertipe data string.
//Fungsi ini tidak memiliki nilai kembali. Ia hanya mencetak "Hello nama" ke layar dengan menggunakan perintah fmt.Println, di mana nama adalah nilai yang diberikan saat fungsi dipanggil.

func LuasPersegi(panjang int, lebar int) int {
	return panjang * lebar
}

//Fungsi ini bernama LuasPersegi dan memiliki dua parameter bernama panjang dan lebar yang bertipe data integer.
//Fungsi ini memiliki nilai kembali bertipe data integer. Nilai kembali dihasilkan dari perkalian dari nilai panjang dan lebar.

func main() {

	//pemanggilan fungsi
	sayHello()
	sayHelloTo("Nursandy Ihksan")
	fmt.Println(LuasPersegi(4, 4))
}
