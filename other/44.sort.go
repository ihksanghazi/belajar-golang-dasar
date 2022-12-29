package main

import (
	"fmt"
	"sort"
)

type user struct {
	name string
	age  int
}

type userSlice []user

func (value userSlice) Len() int {
	return len(value)
}

func (value userSlice) Less(i, j int) bool {
	return value[i].age < value[j].age
}

func (value userSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	users := []user{
		{"Sandy", 21},
		{"Azhi", 20},
		{"Andri", 21},
		{"Sendi", 20},
	}

	sort.Sort(userSlice(users))

	fmt.Println(users)
}
