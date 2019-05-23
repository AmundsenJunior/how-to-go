// Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS “x” and “y” and “z”
// 42
// “James Bond”
// true
// Now print the values stored in those variables using
// a single print statement
// multiple print statements
// code: here’s the solution: https://play.golang.org/p/yYXnWXIQNa

package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println("x:", x, "y:", y, "z:", z)
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)
}
