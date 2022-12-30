package main

import (
	"fmt"
	"reflect"
)

type sample struct {
	Name string `required:"true" max:"10"`
}

func isValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

type contoh struct {
	name        string `required:"true"`
	description string `required:"true"`
}

func main() {
	sample := sample{"Nursandy Ihksan"}
	var sampleType reflect.Type = reflect.TypeOf(sample)

	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))

	fmt.Println(isValid(sample))
	contoh := contoh{"", ""}
	fmt.Println(isValid(contoh))
}
