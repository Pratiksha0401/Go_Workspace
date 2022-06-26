package main

import (
	"fmt"
	"math"
)

type result interface{
	area() float64
	perimeter() float64
}

type rectangle struct {
	length, breadth float64
}

type circle struct{
	radius float64
}

func measure(r result) {
    fmt.Println(r)
    fmt.Println(r.area())
    fmt.Println(r.perimeter())
}
func main() {
    r := rectangle{length: 3, breadth: 4}
    c := circle{radius: 5}
	measure(r)
    measure(c)
}

func printPerimeter(r result) {
	fmt.Println(r.perimeter())
}

func (r rectangle) area() float64 {
	return r.length * r.breadth
}

func (c circle) area() float64 {
	return 2*math.Pi*math.Pow(c.radius, 2)
}

func (r rectangle) perimeter() float64 {
	return (2*r.length) + (2*r.breadth)
}

func (c circle) perimeter() float64 {
	return 2*math.Pi*c.radius
}