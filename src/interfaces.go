package main

import (
	"fmt"
	"math"

)

type Shape interface{
	area() float64
} 

type Circle struct {
	x,y,radius float64
}

type Rectangle struct {
	length, width float64
}

// circle area
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// rectangle area
func (r Rectangle) area() float64 {
	return r.length * r.width
}

// 
func getArea(s Shape) float64 {
	return s.area()
}

func main(){
	
	rect1 := Rectangle{length: 12, width:4}
	circle1 := Circle{x:0, y:0, radius: 7}

	fmt.Println(getArea(rect1))
	fmt.Println(getArea(circle1))
	
}
