package main

import "hash/crc32"

func main()  {
	// create a hasher
	h := crc32.NewIEEE()
	// write data to it
	h.Write([]byte("test"))
}
