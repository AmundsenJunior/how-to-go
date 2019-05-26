/*
  Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
  first name
  last name
  favorite ice cream flavors
  Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.
  solution: https://play.golang.org/p/ouRHmH_POg
*/

package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	flavors   []string
}

func main() {
	p1 := person{
		firstname: "Scott",
		lastname:  "Russell",
		flavors: []string{
			"vanilla",
			"cookies 'n' cream",
		},
	}

	p2 := person{
		firstname: "Penny",
		lastname:  "Nickel",
		flavors: []string{
			"tuna",
			"grass",
		},
	}

	printPerson(p1)
	printPerson(p2)
}

func printPerson(p person) {
	fmt.Println(p.firstname, p.lastname)
	for _, v := range p.flavors {
		fmt.Printf("\t%v\n", v)
	}
}
