package main

import (
	"fmt"
)

func zero(x int){
	x = 0
}

func zeroPtr(xPtr *int){
	*xPtr = 0
}

func addArr(args ...int){
	total := 0
	for i := 0; i < len(args); i++{
		total += args[i]
	}
	fmt.Println(total)
}

func main() {
	x:= 5
	zero(x)
	fmt.Println(x)

	zeroPtr(&x)
	fmt.Println(x)

	addArr(1,2,3,4,5)
}