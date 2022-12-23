package main

import "fmt"

// Menentukan interface bernama Shape
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Menentukan struct bernama Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Menambahkan method Area pada struct Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Menambahkan method Perimeter pada struct Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

func main() {
	// Membuat instance dari struct Rectangle
	r := Rectangle{Width: 10, Height: 5}

	// Memanggil method Area dan Perimeter
	fmt.Println(r.Area())      // Output: 50
	fmt.Println(r.Perimeter()) // Output: 30
}
