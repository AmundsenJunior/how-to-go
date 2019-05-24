/*
  Using the code from the previous exercise,
  At the package level scope, assign the following values to the three variables
  for x assign 42
  for y assign “James Bond”
  for z assign true
  in func main
  use fmt.Sprintf to print all of the VALUES to one single string. 
  ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER “s”
  print out the value stored by variable “s”
  code: here’s the solution: https://play.golang.org/p/QFctSQB_h3 
  */

package main

import "fmt"

var a = 42
var b = "James Bond"
var c = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v", a, b, c)
	fmt.Println(s)
}
