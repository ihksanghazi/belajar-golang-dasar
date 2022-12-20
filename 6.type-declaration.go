package main

import "fmt"

func main() {
	type noKTP string
	type Married bool

	var noKtpSandy noKTP = "123456789"
	var marriedStatus Married = false

	//Pada baris pertama, kita mendefinisikan tipe data baru bernama noKTP yang merupakan tipe data string. Pada baris kedua, kita mendefinisikan tipe data baru bernama Married yang merupakan tipe data boolean.
	//Pada baris keempat, kita mendeklarasikan variabel noKtpSandy dengan tipe data noKTP dan memberinya nilai "123456789". Pada baris kelima, kita mendeklarasikan variabel marriedStatus dengan tipe data Married dan memberinya nilai false.

	fmt.Println(noKtpSandy)
	fmt.Println(marriedStatus)
}
