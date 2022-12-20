package main

import "fmt"

func main() {
	var name string
	name = "Nursandy Ihksan"
	fmt.Println(name)
	name = "Ihksan Ghazi"
	fmt.Println(name)

	//variabel bernama "name" yang memiliki tipe data string. Kemudian, nilai dari variabel "name" diisi dengan string "Nursandy Ihksan" dan dicetak ke layar menggunakan perintah fmt.Println(name). Kemudian, nilai dari variabel "name" diubah menjadi "Ihksan Ghazi" dan dicetak ke layar lagi menggunakan perintah fmt.Println(name).

	var name = "Nursandy Ihksan" // langsung tahu kalo itu string
	fmt.Println(name)

	country := "Indonesia" // bentuk penulisan singkat variabel
	fmt.Println(country)
	handphone := 68
	fmt.Println(handphone)

	var (
		fistname = "Nursandy"
		lastname = "Ihksan"
	)

	fmt.Println(fistname)
	fmt.Println(lastname)

	//variabel "firstname" dan "lastname" dideklarasikan dan diinisialisasi sekaligus menggunakan sintaks "var ( ... )". Sintaks ini dapat digunakan untuk memperkenalkan lebih dari satu variabel sekaligus dalam satu baris kode
}
