/*
  create a type SQUARE
  create a type CIRCLE
  attach a method to each that calculates AREA and returns it
  circle area= π r 2
  square area = L * W
  create a type SHAPE that defines an interface as anything that has the AREA method
  create a func INFO which takes type shape and then prints the area
  create a value of type square
  create a value of type circle
  use func info to print the area of square
  use func info to print the area of circle
  code: https://play.golang.org/p/NGGikHNvHv
*/

package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	a := math.Exp2(s.length)
	return a
}

func (c circle) area() float64 {
	a := math.Pi * math.Exp2(c.radius)
	return a
}

func info(s shape) {
	area := s.area()
	switch s.(type) {
	case square:
		fmt.Printf("The area of a square is %v\n", area)
	case circle:
		fmt.Printf("The area of a circle is %v\n", area)
	}
}

func main() {
	square1 := square{
		length: 15.5,
	}

	circle1 := circle{
		radius: 3.333,
	}

	info(square1)
	info(circle1)
}
