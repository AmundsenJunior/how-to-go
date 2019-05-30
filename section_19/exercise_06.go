// creating custom Sort functions for any data structure using sort.Interface
// sort.Interface is an interface needing the Len, Less, and Swap methods for any type
// with those defined, can then execute sort.Sort on the collection

package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Last  string
	Age   int
}

type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type ByFirst []Person

func (f ByFirst) Len() int {
	return len(f)
}

func (f ByFirst) Less(i, j int) bool {
	return f[i].First < f[j].First
}

func (f ByFirst) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func main() {
	p1 := Person{
		First: "Scott",
		Last: "Russell",
		Age: 36,
	}

	p2 := Person{
		First: "Penny",
		Last: "Nickel",
		Age: 60,
	}

	people := []Person{p1, p2,}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

	sort.Sort(ByFirst(people))
	fmt.Println(people)
}

