/*
  Using iota, create 4 constants for the NEXT 4 years. Print the constant values.
  solution: https://play.golang.org/p/MDLF3v5EGT
*/

package main

import "fmt"

const (
	year = 2019
	a    = year + iota
	b    = year + iota
	c    = year + iota
	d    = year + iota
)

func main() {
	fmt.Println(a, b, c, d)
}
