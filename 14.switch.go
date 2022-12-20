package main

import "fmt"

func main() {
	var name = "sandy"

	switch name {
	case "sandy":
		fmt.Println("Hello Sandy")
	case "Nur":
		fmt.Println("Hello Nur")
	case "Ihksan":
		fmt.Println("hello Ihksan")
	default:
		fmt.Println("Hello")
	}

	// Pada baris pertama, variabel name diberi nilai "sandy". Kemudian, pada baris kedua terdapat percabangan switch yang akan mengevaluasi nilai variabel name.
	// Pada baris ketiga sampai kelima, terdapat tiga percabangan case yang akan diuji apakah sama dengan nilai dari variabel name. Jika salah satu dari kondisi tersebut terpenuhi, maka akan mencetak string yang sesuai ke layar.
	// Jika tidak ada kondisi yang terpenuhi, maka akan diikuti oleh percabangan default pada baris keenam, yang akan mencetak string "Hello" ke layar.

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama sudah benar")
	default:
		fmt.Println("Tidak Keduanya")
	}

	// Pada baris pertama, terdapat percabangan switch yang mengevaluasi panjang string dari variabel name. Panjang string tersebut dihitung menggunakan fungsi len(), dan hasilnya disimpan pada variabel length. Kemudian, variabel length akan diuji apakah lebih besar dari 5.
	// Pada baris kedua dan ketiga, terdapat dua percabangan case yang akan diuji apakah sama dengan nilai dari percabangan switch pertama. Jika salah satu dari kondisi tersebut terpenuhi, maka akan mencetak string yang sesuai ke layar.
	// Jika tidak ada kondisi yang terpenuhi, maka akan diikuti oleh percabangan default pada baris keempat, yang akan mencetak string "Tidak Keduanya" ke layar.

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("a")
	case length > 5:
		fmt.Println("b")
	default:
		fmt.Println("c")
	}

	//Pada baris pertama, variabel length diberi nilai sesuai dengan panjang string dari variabel name. Kemudian, pada baris kedua terdapat percabangan switch yang tidak mengevaluasi nilai apapun.
	// Pada baris ketiga sampai kelima, terdapat tiga percabangan case yang akan diuji apakah kondisinya terpenuhi. Jika salah satu dari kondisi tersebut terpenuhi, maka akan mencetak string yang sesuai ke layar.
	// Jika tidak ada kondisi yang terpenuhi, maka akan diikuti oleh percabangan default pada baris keempat, yang akan mencetak string "c" ke layar.

}
