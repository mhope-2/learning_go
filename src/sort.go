package main

import (
	"fmt"
	"sort"
)

type APerson struct{
	Name string
	Age int
}

type ByName []APerson

// By Name
func (ps ByName) Len() int {
	return len(ps)
}

func (ps ByName) Less(i, j int) bool {
	return ps[i].Name > ps[j].Name
}

func (ps ByName) Swap(i, j int)  {
	ps[i], ps[j] = ps[j], ps[i]
}

// By Age
type ByAge []APerson
func (this ByAge) Len() int {
	return len(this)
}

func (this ByAge) Less(i, j int) bool {
	return this[i].Age < this[j].Age
}
func (this ByAge) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func main(){
	kids := []APerson{
		{"Jack", 9},
		{"Jill", 10},
	}

	sort.Sort(ByName(kids))
	fmt.Println(kids)
}