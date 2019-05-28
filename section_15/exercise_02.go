/*
  create a func with the identifier foo that
  takes in a variadic parameter of type int
  pass in a value of type []int into your func (unfurl the []int)
  returns the sum of all values of type int passed in
  create a func with the identifier bar that
  takes in a parameter of type []int
  returns the sum of all values of type int passed in
  code: https://play.golang.org/p/B0yRxtBQPD
*/

package main

import "fmt"

func main() {
	ints := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55,}
	s1 := woo(ints...)
	s2 := baz(ints)
	fmt.Printf("sum from woo is %d\nsum from baz is %d\n", s1, s2)
}

func woo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}

	return sum

}

func baz(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}

	return sum

}
