package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Nur")
	data.PushBack("Sandy")
	data.PushBack("Ihksan")
	data.PushFront("Azhi")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	// for element := data.Front(); element != nil; element = element.Next() {
	// 	fmt.Println(element.Value)
	// }

	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
