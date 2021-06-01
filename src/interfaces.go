package main
import (
	"fmt";
)

type Shape interface {
	area() float64
}


func totalArea(shapes ...Shape) float64 { 
	var area float64 =0.0
	for _, s := range shapes { 
		area += s.area()
	}
	return area 
}

func main(){

	fmt.Println("")
}