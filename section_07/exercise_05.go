/*
  Create a variable of type string using a raw string literal. Print it.
  solution: https://play.golang.org/p/dLy36A-V-w
*/

package main

import "fmt"

func main() {
	s := `this is
not a single-
	line
		string`

	fmt.Println(s)
}
