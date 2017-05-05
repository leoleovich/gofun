package main

import (
    "fmt"
    "math"
)

type Shape interface {
    area() float64
}

type Rectangle struct {
    Width, Height float64
}

func (rectangle Rectangle) area() float64 {
    return rectangle.Width*rectangle.Height
}

type Circle struct {
    Radius float64
}

func (circle Circle) area() float64 {
    return math.Pi*math.Pow(circle.Radius, 2.0)
}

func getAria(shape Shape) float64 {
    return shape.area()
}

func main() {
    r := Rectangle{2,4}
    c := Circle{3}

    fmt.Println(getAria(r))
    fmt.Println(getAria(c))
}
