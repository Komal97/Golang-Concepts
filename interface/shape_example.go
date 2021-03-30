package main

import (
	"fmt"
	"math"
)

/**
Interface Shape
Circle implements Shape
Rectangle implements Shape

Both implements internally because both has all the functions defined in Shape interface
*/

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	width float64
	height float64
}

func (r rectangle) area() float64{
	return r.width * r.height
}

func (c circle) area() float64{
	return math.Pi * c.radius * c.radius
}

func implementShape(){

	fmt.Println("--------------- SHAPE -----------------")

	c1 := circle{radius:4.5}
	r1 := rectangle{width: 5, height: 7}

	shapes := []shape{c1, r1}
	fmt.Println(shapes[0].area())
	fmt.Println(shapes[1].area())
}