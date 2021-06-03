package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main()  {
	// Define flags
	maxp := flag.Int("max", 100, "the max value")
	// Parse
	flag.Parse()
	// Generate a number between 0 and max
	fmt.Println(rand.Intn(*maxp))
	fmt.Println(flag.Args())
}
