package main

import (
	"fmt"
)

func main() {
	var x [5]float64
        x[0] = 98
        x[1] = 93
        x[2] = 77
        x[3] = 82
        x[4] = 83
    
	var total float64 = 0 
	for i:=0; i<5; i++{
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))

	// Array
	y := [5]float64{ 
		98,
        93,
        77,
        82,
        83,
	}
	fmt.Println("The length of y is", len(y))

	// Slice

	the_slice := make([]float64, 5)
	fmt.Println("Length of the_slice:", len(the_slice))

	arr := [5]float64{
		1,2,3,4,5,
	}
	slice_2 := arr[0:5]
	fmt.Println(slice_2)

	fmt.Println()

	slice1 := []int{1,2,3}
	slice2 := append(slice1, 4, 5) 
	fmt.Println(slice1, slice2)
}