package main

import (
	"fmt"; 
	"math"
)

// struct definition
type Circle struct { 
	x,y,r float64
}

func circleArea(c Circle) float64 { 
	return math.Pi * c.r * c.r
}

func main() {

	// struct instance (create space in memory)
	// var c Circle

	// assign values
	c := Circle{ x: 0, y: 0, r: 5 } 

	// pointer to the struct
	// cPtr := &Circle{0, 0, 5}
	// fmt.Println(circleArea(&c))

	// get struct variable values
	fmt.Println(c.x, c.y, c.r)

	// print area of circle
	fmt.Println(circleArea(c))
}


