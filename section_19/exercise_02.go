/*
  Starting with this code (https://play.golang.org/p/b_UuCcZag9),
  unmarshal the JSON into a Go data structure.
  Hint: https://mholt.github.io/json-to-go/
  code:
  code setup (just fyi, not needed for exercise):
  https://play.golang.org/p/nWPP5Z2Q4e
  solution:
  https://play.golang.org/p/r8oSG8DIPR
*/

package main

import (
	"encoding/json"
	"fmt"
)

type character struct {
	First string `json:First`
	Last string `json:Last`
	Age int `json:Age`
	Sayings []string `json:Sayings`
}

func main() {
	s := `[
  {
    "First":"James",
    "Last":"Bond",
    "Age":32,
    "Sayings":[
      "Shaken, not stirred",
      "Youth is no guarantee of innovation",
      "In his majesty's royal service"
    ]
  },{
    "First":"Miss",
    "Last":"Moneypenny",
    "Age":27,
    "Sayings":[
      "James, it is soo good to see you",
      "Would you like me to take care of that for you, James?",
      "I would really prefer to be a secret agent myself."
    ]
  },{
    "First":"M",
    "Last":"Hmmmm",
    "Age":54,
    "Sayings":[
      "Oh, James. You didn't.",
      "Dear God, what has James done now?",
      "Can someone please tell me where James Bond is?"
    ]
  }
]`

	fmt.Println(s)

	var characters []character

	err := json.Unmarshal([]byte(s), &characters)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range characters {
		fmt.Printf("\n%s %s\t%d\n", v.First, v.Last, v.Age)
		for _, s := range v.Sayings {
			fmt.Printf("\t%s\n", s)
		}
	}
}
