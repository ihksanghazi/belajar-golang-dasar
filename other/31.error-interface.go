package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian Tidak Boleh Nol")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	// var contohError error = errors.New("ups Error")

	hasil, err := pembagian(10, 0)
	if err == nil {
		fmt.Println("Hasil = ", hasil)
	} else {
		fmt.Println("Error = ", err.Error())
	}
}
