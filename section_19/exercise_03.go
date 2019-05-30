/*
  Starting with this code (https://play.golang.org/p/BVRZTdlUZ_),
  encode to JSON the []user sending the results to Stdout.
  Hint: you will need to use json.NewEncoder(os.Stdout).encode(v interface{})
  solution: https://play.golang.org/p/ql90D1OQqd
*/

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := User{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := User{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := User{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []User{u1, u2, u3}

	output := json.NewEncoder(os.Stdout)
	output.SetIndent("", "	") // pretty-print your Json by putting a tab or other characters in for default "" indent

	err := output.Encode(users)
	if err != nil {
		fmt.Println(err)
		return
	}
}
