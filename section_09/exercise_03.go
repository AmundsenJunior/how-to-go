/*
  Create a for loop using this syntax
  for condition { }
  Have it print out the years you have been alive.
  solution: https://play.golang.org/p/tnyqBPJ-i5
*/

package main

import "fmt"

func main() {
	year := 1982
	for year <= 2019 {
		fmt.Println(year)
		year++
	}
}
