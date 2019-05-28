/*
  Closure is when we have “enclosed” the scope of a variable in some code block.
  For this hands-on exercise, create a func which “encloses” the scope of a variable:
  code: https://play.golang.org/p/a56uWtoFSL
*/

package main

import "fmt"

func enclosure() {
	var x int
	fmt.Println("var x int within func enclosure has value", x)
	x++
}
func main() {
	x := 10
	fmt.Println("var x int within func main has value", x, "before calling func enclosure")

	enclosure()

	fmt.Println("var x int within func main has value", x, "after calling func enclosure")
}


