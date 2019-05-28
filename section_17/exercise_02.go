/*
  create a person struct
  create a func called “changeMe” with a *person as a parameter
    change the value stored at the *person address
      important
        to dereference a struct, use (*value).field
          p1.first
          (*p1).first
        “As an exception, if the type of x is a named pointer type
        and (*x).f is a valid selector expression denoting a field (but not a method),
        x.f is shorthand for (*x).f.”
          https://golang.org/ref/spec#Selectors
  create a value of type person
    print out the value
  in func main
    call “changeMe”
  in func main
    print out the value
  code: https://play.golang.org/p/AehM662HKS
*/

package main

import "fmt"

type person struct {
	first string
	last string
}

func (p *person) getPerson() (string, string) {
	return p.first, p.last // or return (*p).first, (*p).last
}

// pointer receiver changes the value of type person
func (p *person) setPerson(first string, last string) {
	p.first = first // or (*p).first
	p.last = last // or (*p).last
}

/* nonpointer receiver does not change the existing value
func (p person) setPerson(first string, last string) {
	p.first = first
	p.last = last
}*/

func main() {
	p1 := person{
		first: "Scott",
		last: "Russell",
	}
	fmt.Println(p1.getPerson())

	p1.setPerson("Penny", "Nickel")

	fmt.Println(p1.getPerson())
}

