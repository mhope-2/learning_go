package main
import "fmt"

var x string = "Expert"

func main(){
	 fmt.Println("Hello my name is", "Mke")
	 fmt.Println("1 + 1 =", 1.0 + 1.0)
	 fmt.Println(len("Hello"))
	 fmt.Println("Hello"[1])
	 fmt.Println(!true)

	 num := 7
	 fmt.Println(num)

	 const name string = "Kay"

	 // declaring multiple varibles

	// var ( 
	// 	a = 5
	// 	b = 10
	// 	c = 15 
	// )

	for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}

	i := 4

	switch i {
		case 0: fmt.Println("Zero")
		case 1: fmt.Println("One")
		case 2: fmt.Println("Two")
		case 3: fmt.Println("Three")
		case 4: fmt.Println("Four")
	default: fmt.Println("Unknown")
	}

	fmt.Println()

}

func average(xs []float64) float64 { 
	panic("Not Implemented")
}

