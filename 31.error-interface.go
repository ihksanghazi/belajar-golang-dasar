package main

import "fmt"

func main() {

	var myInterface interface{}
	fmt.Println(myInterface.SomeMethod()) // akan menyebabkan error "Nil pointer dereference"

	var myInterface1 interface{}
	var str string = myInterface1.(string) // akan menyebabkan error "Type assertion error"

	var myInterface2 interface{} = 123
	var str2 string = myInterface2.(string) // akan menyebabkan error "Invalid type assertion"

}
