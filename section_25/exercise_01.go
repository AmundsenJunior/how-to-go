/*
  Start with this code (https://play.golang.org/p/3W69TH4nON).
  Instead of using the blank identifier, make sure the code is checking and handling the error.
  solution:
  https://play.golang.org/p/tn8oiuL1Yn
*/

package main

import (
	"encoding/json"
	"fmt"
)

type agent struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	a1 := agent{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(a1)
	if err != nil {
		fmt.Println("Error encountered marshalling JSON:", err)
		return
		// or log.Fatalf("Error encountered marshalling JSON: %v\n", err)
		// or panic(err)
	}
	fmt.Println(string(bs))

}
