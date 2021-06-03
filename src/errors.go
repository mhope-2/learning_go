package main
import (
	"errors"
	"fmt"
)

func main() {
	//custom error msg
	err := errors.New("Custom Error Message")
	fmt.Println(err)
}