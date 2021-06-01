package main
import (
	"fmt"
	"strings"
	"os"
	"io/ioutil"
	"path/filepath"
)

func main() {
	// string pkg
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

	// convert a string to a slice of bytes 
	arr := []byte("test")
	fmt.Println(arr)

	// convert slice of bytes to string
	str := string([]byte{'t','e','s','t'})	
	fmt.Println(str)


	/* FILE READ USING the os PACKAGE*/
	// os pkg
	file, err := os.Open("1.txt")
	if err != nil {
		// handle errors
		fmt.Println("Error occurred")
	}
	fmt.Println(file)
	defer file.Close()

	// get the file size
	stat, err := file.Stat() 
	if err != nil {
		return
	}
	// char count
	fmt.Println(stat.Size())

	// read the file
	bs := make([]byte, stat.Size()) 
	_, err = file.Read(bs) 
	if err != nil {
		return
	}

	fmt.Println("File content (from os)")
	file_str := string(bs)
	fmt.Println(file_str)

	fmt.Println()


	/* FILE READ USING io/util package */
	bs, err = ioutil.ReadFile("1.txt")
	if err != nil {
		return
	}
	io_file_str := string(bs)
	fmt.Println("File content (from ioutil)")
	fmt.Println(io_file_str)


	/* CREATE FILE */
	newFile, err := os.Create("created.txt") 
	if err != nil {
		// handle the error here
		return
	}

	defer newFile.Close()
	newFile.WriteString("This is a file created with os.Create()")


	/* READ DIR CONTENT */
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1) 
	if err != nil {
		return
		}
	
	for _, fi := range fileInfos {
		fmt.Println(fi.Name()) 
	}

	// Walk the dir and sub dirs
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
		})

}




























































