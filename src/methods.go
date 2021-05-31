package main

import ( "fmt";
		 "math"
		)

// circle struct		
type Circle struct {
	x,y,r float64
}

// circle area method
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}


// person struct
type Person struct{
	name string
}

type Android struct { 
	Person
	Model string 
}

// person talk method
func (p *Person) talk() {
	fmt.Println("My name is", p.name)
}

func main(){

	c := Circle{ x: 0, y: 0, r: 5 } 
	fmt.Println(c.area())

	p := Person{name: "Hope"}
	p.talk()

}