/*
  Create a user defined struct with
  the identifier “person”
  the fields:
  first
  last
  age
  attach a method to type person with
  the identifier “speak”
  the method should have the person say their name and age
  create a value of type person
  call the method from the value of type person
  code: https://play.golang.org/p/NnXyWdqbbw
*/

package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func (p person) speak() {
	fmt.Printf("Hey y'all, my name is %s %s and I'm %d years old.\n", p.first, p.last, p.age)
}

func main() {
	p1 := person{
		first: "Scott",
		last: "Russell",
		age: 36,
	}

	p1.speak()
}

