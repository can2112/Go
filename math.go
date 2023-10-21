package main

import (
	"fmt"
	"math"
)

type Sizer interface {
	area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) area() float64 {
	return c.Radius * c.Radius * math.Pi
}

type Square struct {
	Side float64
}

func (c Square) area() float64 {
	return c.Side * c.Side
}

func Less(a Sizer, b Sizer) Sizer {
	if a.area() < b.area() {
		return a
	}
	return b
}

func mathService() {
	fmt.Println("hello world!")
	c1, c2 := Circle{Radius: 2.5}, Circle{Radius: 5.0}
	sq1, sq2 := Square{Side: 3.0}, Square{Side: 4.0}
	fmt.Println(Less(c1, c2))
	fmt.Println(Less(sq1, sq2))
}
