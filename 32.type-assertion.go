package main

import "fmt"

func random() interface{} {
	return "OK"
}

func main() {
	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	// resultInt := result.(int) // panic
	// fmt.Println(resultInt)

	switch value := result.(type) {
	case string:
		fmt.Println("Value ", value, " is String")
	case bool:
		fmt.Println("Value ", value, " is Boolean")
	case int:
		fmt.Println("Value ", value, " is Integer")
	default:
		fmt.Println("Uknown !!!")
	}

}
