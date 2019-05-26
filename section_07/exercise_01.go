/*
  Write a program that prints a number in decimal, binary, and hex
  solution: https://play.golang.org/p/bAQxcEuK8O
*/

package main

import "fmt"

func main() {
	x := 112358
	fmt.Printf("%d\t%b\t%x\n", x, x, x)
}
