/*
  Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
  code: https://play.golang.org/p/XluEuUD0Nw
*/

package main

import "fmt"

func main() {
	fmt.Println("first things first...")

	defer deferred()

	fmt.Println("this is coded last but...")
}

func deferred() {
	fmt.Println("this runs at the end of the calling function.")
}
