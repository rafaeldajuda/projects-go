package main

import (
	"fmt"
	"math"
)

// Interface Shape com as funções Area e Perimeter
type Shape interface {
	Area() float64
	Perimeter() float64
	Nome() string
}

// Estrutura Circle que implementa a interface Shape
type Circle struct {
	Ratio float64
}

// Cálculo da área de um círculo
func (c Circle) Area() float64 {
	return math.Pi * c.Ratio * c.Ratio
}

// Cálculo do perímetro de um círculo
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Ratio
}

// Estrutura Rectangle que implementa a interface Shape
type Rectangle struct {
	Width  float64
	Height float64
}

// Cálculo da área de um retângulo
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Cálculo do perímetro de um retângulo
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Mostra a área e o perímetro de um Shape
func GetShapeValues(shape Shape) {
	fmt.Printf("Area: %v\nPerimeter: %v\n", shape.Area(), shape.Perimeter())
}

func main() {
	fmt.Println("Circle")
	c := Circle{Ratio: 2}
	GetShapeValues(c)

	fmt.Println("Rectangle")
	r := Rectangle{Width: 2, Height: 2.5}
	GetShapeValues(r)
}
