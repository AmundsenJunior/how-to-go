/*
  Write a program that
  assigns an int to a variable
  prints that int in decimal, binary, and hex
  shifts the bits of that int over 1 position to the left, and assigns that to a variable
  prints that variable in decimal, binary, and hex
  solution: https://play.golang.org/p/Ms964T8SbH
 */

package main

import "fmt"

func main() {
	x := 2357
	fmt.Printf("%d\t%b\t%x\n", x, x, x)
	y := x << 1
	fmt.Printf("%d\t%b\t%x\n", y, y, y)
}