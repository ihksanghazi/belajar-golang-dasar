package main

import "fmt"

func main() {
	var name [3]string

	name[0] = "Nur"
	name[1] = "Sandy"
	name[2] = "Ihksan"

	//pertama-tama terdapat deklarasi array name yang terdiri dari 3 elemen bertipe string. Kemudian, nilai dari masing-masing elemen dari array tersebut diisi dengan "Nur", "Sandy", dan "Ihksan".

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	var values = [3]int{
		1, 2, 3,
	}

	//Deklarasi array dengan sintaks shorthand seperti ini akan secara otomatis menentukan panjang dari array sesuai dengan jumlah elemen yang dideklarasikan. Sebagai contoh, dalam kode di atas, panjang dari array values akan secara otomatis ditentukan menjadi 3 karena terdapat 3 elemen yang dideklarasikan.

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println("Jumlah array name = ", len(name))
	fmt.Println("Jumlah array values = ", len(values))
}
