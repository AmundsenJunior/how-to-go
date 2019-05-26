/*
  Using a COMPOSITE LITERAL:
  create an ARRAY which holds 5 VALUES of TYPE int
  assign VALUES to each index position.
  Range over the array and print the values out.
  Using format printing
  solution: https://play.golang.org/p/tD0SzV3hdf
*/

package main

import "fmt"

func main() {
	x := [5]int{1, 1, 2, 3, 5}
	for _, v := range x {
		fmt.Println(v)
	}
}

