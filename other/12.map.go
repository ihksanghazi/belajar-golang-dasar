package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "eko",
		"address": "Jakarta",
	}

	person["title"] = "Programmer"

	//Dalam kode di atas, terdapat deklarasi variabel person yang merupakan sebuah map yang menyimpan pasangan-pasangan key-value bertipe string. Map person terdiri dari 2 pasangan key-value yang dideklarasikan secara langsung, yaitu "name" dengan value "eko" dan "address" dengan value "Jakarta".
	//Kemudian, terdapat baris kode yang menambah pasangan key-value baru ke dalam map person, yaitu "title" dengan value "Programmer". Dengan demikian, map person akan terdiri dari 3 pasangan key-value, yaitu "name" dengan value "eko", "address" dengan value "Jakarta", dan "title" dengan value "Programmer".

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Nursandy Ihksan"
	book["ups"] = "Salah"

	//Dalam kode di atas, terdapat deklarasi variabel book yang merupakan sebuah map yang menyimpan pasangan-pasangan key-value bertipe string. Variabel book dideklarasikan dengan menggunakan syntax var book map[string]string = make(map[string]string). Syntax ini digunakan untuk mengalokasikan memori untuk map book dengan tipe data key dan value yang telah ditentukan, yaitu string.
	//Setelah itu, terdapat 3 baris kode yang menambah pasangan key-value ke dalam map book. Pertama, baris kode yang menambah pasangan key-value "title" dengan value "Belajar Golang". Kemudian, baris kode yang menambah pasangan key-value "author" dengan value "Nursandy Ihksan". Dan terakhir, baris kode yang menambah pasangan key-value "ups" dengan value "Salah".

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	//Dalam kode di atas, terdapat variabel book yang merupakan sebuah map yang menyimpan pasangan-pasangan key-value bertipe string. Kemudian, terdapat pemanggilan fungsi delete() yang menerima 2 argument, yaitu book dan string "ups". Fungsi delete() akan mencari pasangan key-value dengan key "ups" dalam map book dan menghapusnya.
	//Jika pasangan key-value dengan key "ups" ditemukan dalam map book, maka pasangan tersebut akan dihapus dan fungsi delete() akan mengembalikan nilai true. Jika pasangan key-value dengan key "ups" tidak ditemukan dalam map book, maka fungsi delete() akan mengembalikan nilai false.

	fmt.Println(book)
	fmt.Println(len(book))
}
