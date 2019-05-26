/*
  Using a COMPOSITE LITERAL:
  create a SLICE of TYPE int
  assign 10 VALUES
  Range over the slice and print the values out.
  Using format printing
  print out the TYPE of the slice
  solution: https://play.golang.org/p/sAQeFB7DIs
 */

package main

import "fmt"

func main()  {
	x := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}

	fmt.Printf("slice is of type %T\n", x)

	for i, v := range x {
		fmt.Printf("%v, %v\n", i, v)
	}
}
