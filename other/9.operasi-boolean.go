package main

import "fmt"

func main() {
	var ujian = 80
	var absensi = 80

	// var lulusUjian = ujian >= 80
	// var lulusAbsensi = absensi >= 80
	// fmt.Println(lulusUjian)
	// fmt.Println(lulusAbsensi)

	// var lulus = lulusUjian && lulusAbsensi

	// fmt.Println(lulus)

	fmt.Println(ujian >= 80 && absensi >= 80)

	//Ekspresi boolean tersebut mengevaluasi apakah nilai dari variabel ujian dan absensi masing-masing lebih besar sama dengan 80. Jika kedua ekspresi tersebut bernilai benar (true), maka ekspresi boolean ujian >= 80 && absensi >= 80 akan bernilai true dan akan dicetak ke layar. Sebaliknya, jika salah satu atau kedua ekspresi tersebut bernilai salah (false), maka ekspresi boolean tersebut akan bernilai false dan tidak akan dicetak ke layar.
}
