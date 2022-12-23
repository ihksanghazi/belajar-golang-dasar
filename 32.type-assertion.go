package main

import "fmt"

func main() {
	var myInterface interface{} = 123

	// Type assertion dengan syntax "x.(T)"
	var i int = myInterface.(int)
	fmt.Println(i) // Output: 123

	// Type assertion dengan syntax "x.(T), ok"
	var f float64
	var ok bool
	f, ok = myInterface.(float64)
	fmt.Println(f, ok) // Output: 0 false
}
