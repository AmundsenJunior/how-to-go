/*
  Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:

  "James", "Bond", "Shaken, not stirred"
  "Miss", "Moneypenny", "Helloooooo, James."

  Range over the records, then range over the data in each record.
  solution: https://play.golang.org/p/FMM4c2PhZg
*/

package main

import "fmt"

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James"}

	z := [][]string{x, y}

	for _, u := range z {
		for _, v := range u {
			fmt.Println(v)
		}
	}
}
