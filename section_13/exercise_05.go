/*
  can't assign to promoted fields of an embedded struct
  can only call promoted fields
*/

package main

import (
	"fmt"
)

type human struct {
	first, last string
}

type engineer struct {
	human
	licensed bool
}

func main() {
	p1 := engineer{
		human: human{
			first: "scott",
			last:  "russell",
		},
		licensed: true,
	}

	// p2 := engineer{
	// 	first:    "penny",
	// 	last:     "nickel",
	// 	licensed: false,
	// }

	fmt.Println(p1.first, p1.last, p1.licensed)
}
