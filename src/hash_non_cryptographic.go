package main

import (
	"fmt"
	"hash/crc32"
)

func hasherInit(){
	// create a hasher
	h := crc32.NewIEEE()
	// write data to it
	h.Write([]byte("test"))
	// calculate the crc32 checksum
	v := h.Sum32()
	fmt.Println(v)
}

func main()  {

}
