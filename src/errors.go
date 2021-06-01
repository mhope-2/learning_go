package main
import (
	"fmt"
	"errors"
)

func main() {
	err := errors.New("Custom Error Message")
	fmt.Println(err)
}