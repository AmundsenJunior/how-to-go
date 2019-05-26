/*
  Create TYPED and UNTYPED constants. Print the values of the constants.
  solution: https://play.golang.org/p/NutvJXWUx2
*/

package main

import "fmt"

const (
	j int8 = 7
	k      = 6
)

func main() {
	fmt.Println(j)
	fmt.Println(k)
	fmt.Printf("%T\n", j)
	fmt.Printf("%T\n", k)
}
