package main
import (
	"fmt";
	"strings"
)

func main() {
	// substring contained in larger string
	fmt.Println(strings.Contains("test", "es"))

	// count of appearance
	fmt.Println(strings.Count("test", "t"))

	// func HasPrefix(s, prefix string) bool
	fmt.Println(strings.HasPrefix("test", "te"))

	// func HasSuffix(s, suffix string) bool 
	fmt.Println(strings.HasSuffix("test", "st"))

	// func Index(s, sep string) int 
	fmt.Println(strings.Index("test", "s"))

	// func Join(a []string, sep string) string 
	fmt.Println(strings.Join([]string{"a","b","c"}, ","))

	// func Repeat(s string, count int) string 
	fmt.Println(strings.Repeat("a", 5))

	// func Replace(s, old, new string, n int) string 
	fmt.Println(strings.Replace("aaaa", "a", "b", 2))

	// func Split(s, sep string) []string 
	fmt.Println(strings.Split("a-b-c-d-e", "-"))

	// func ToLower(s string) string
	fmt.Println(strings.ToLower("TEST"))

	// func ToUpper(s string) string
	fmt.Println(strings.ToUpper("test"))

	// convert a string to a slice of bytes (and vice versa)
	arr := []byte("test")
	str := string([]byte{'t','e','s','t'})	

	fmt.Println(arr)
	fmt.Println(str)
}