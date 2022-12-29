package main

import (
	"fmt"
	"strconv"
)

func main() {
	bool, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(bool)
	} else {
		fmt.Println(err.Error())
	}

	number, err := strconv.ParseInt("1000000", 10, 64)

	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	value, _ := strconv.Atoi("1453")
	fmt.Println(value)
}
