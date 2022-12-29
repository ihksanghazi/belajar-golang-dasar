package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Nursandy Ihksan", "Ihksan")) // ? string mengandung ihksan , YES, TRUE
	fmt.Println(strings.Contains("Nursandy Ihksan", "Azhi"))   // ? string mengandung Azhi , NO , FAlSE

	fmt.Println(strings.Split("Nursandy Ihksan", " "))

	fmt.Println(strings.ToLower("Nursandy Ihksan"))

	fmt.Println(strings.ToUpper("Nursandy Ihksan"))

	fmt.Println(strings.ToTitle("Nursandy Ihksan"))

	fmt.Println(strings.Trim("			Nursandy Ihksan			", "	"))

	fmt.Println(strings.ReplaceAll("Cahaya Sandy Ihksan", "Cahaya", "Nur"))

}
