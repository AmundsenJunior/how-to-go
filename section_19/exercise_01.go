/*
  Starting with this code (https://play.golang.org/p/_fQUGm9Utvl),
  marshal the []user to JSON.
  There is a little bit of a curve ball in this hands-on exercise -
  remember to ask yourself what you need to do to EXPORT a value from a package.
  solution: https://play.golang.org/p/8BK6PXj3aG
*/

package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Age   int
}

func main() {
	u1 := person{
		First: "James",
		Age:   32,
	}

	u2 := person{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := person{
		First: "M",
		Age:   54,
	}

	people := []person{u1, u2, u3}

	fmt.Println(people)

	jsonString, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonString))
}
