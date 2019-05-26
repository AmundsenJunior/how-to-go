/*
  Create a for loop using this syntax
  for { }
  Have it print out the years you have been alive.
  solution: https://play.golang.org/p/9VpnB-I1Pz
*/

package main

import "fmt"

func main() {
	year := 1982

	for {
		if year > 2019 {
			break
		}
		fmt.Println(year)
		year++
	}
}
