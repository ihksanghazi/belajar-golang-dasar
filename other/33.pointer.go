package main

import "fmt"

type addres struct {
	city, province, country string
}

func changeCountryToIndonesia(a *addres) {
	a.country = "Indonesia"
}

func main() {
	// pass by value
	// address1 := addres{"Subang", "Jawa Barat", "Indonesia"}
	// address2 := address1

	// address2.city = "Bandung"

	// fmt.Println(address1)
	// fmt.Println(address2)

	// pass by references
	var address1 addres = addres{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *addres = &address1
	var address3 *addres = &address1

	address2.city = "Bandung"

	*address2 = addres{"Bandung", "DKI Jakarta", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *addres = new(addres)
	address4.city = "Jakarta"
	fmt.Println(address4)

	// pointer di function

	var alamat = addres{
		city:     "subang",
		province: "Jawa Barat",
		country:  "",
	}

	changeCountryToIndonesia(&alamat)
	fmt.Println(alamat)

}
