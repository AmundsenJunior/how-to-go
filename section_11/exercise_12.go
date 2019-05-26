/*
  multi-dimensional slices don't have to have the same length in each sub-slice,
  so can't necessarily be treated the same as a matrix
 */

package main

import (
	"fmt"
)

func main() {
	x := []string{"scott", "edward", "russell"}
	y := []string{"penny", "nickel", "quarter", "dime"}

	z := [][]string{x, y}
	fmt.Println(z)
}
