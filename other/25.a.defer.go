package main

import "fmt"

func logging() {
	fmt.Println("Selesai Memanggil Function")
}

func runAplication(i int) {
	defer logging()
	fmt.Println("run aplication")
	result := 10 / i
	fmt.Println("result = ", result)
}

func main() {
	runAplication("hahha")
}
