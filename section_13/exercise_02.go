/*
  Take the code from the previous exercise,
  then store the values of type person in a map with the key of last name.
  Access each value in the map. Print out the values, ranging over the slice.
  solution: https://play.golang.org/p/RvvLyAMvGo
*/

package main

import "fmt"

type individual struct {
	firstname string
	lastname  string
	flavors   []string
}

func main() {
	i1 := individual{
		firstname: "Scott",
		lastname:  "Russell",
		flavors: []string{
			"vanilla",
			"cookies 'n' cream",
		},
	}

	i2 := individual{
		firstname: "Penny",
		lastname:  "Nickel",
		flavors: []string{
			"tuna",
			"grass",
		},
	}

	individuals := map[string]individual{
		i1.lastname: i1,
		i2.lastname: i2,
	}

	for k := range individuals {
		printIndividual(individuals[k])
	}
}

func printIndividual(i individual) {
	fmt.Println(i.firstname, i.lastname)
	for _, v := range i.flavors {
		fmt.Printf("\t%v\n", v)
	}
}
