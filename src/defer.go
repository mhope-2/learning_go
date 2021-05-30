package main
import ( "fmt";
		 "os"
		)


func first() { 
	fmt.Println("1st")
}

func second() { 
	fmt.Println("2nd")
}

func third() { 
	fmt.Println("3rd")
}

func main() { 
	second()
	first() 
	third()
	f, _ := os.Open("1.txt") 
	defer f.Close()
	fmt.Println(f)
}