package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	utc := time.Date(2022, 12, 29, 20, 13, 5, 10, time.UTC)

	fmt.Println(utc)

	layout := "2004-01-02"
	parse, _ := time.Parse(layout, "1998-03-29")
	fmt.Println("parse = ", parse)
}
