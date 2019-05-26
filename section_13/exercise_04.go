/*
  Create and use an anonymous struct.
  solution: https://play.golang.org/p/otBHFs-lPp
*/

package main

import "fmt"

func main() {
	anon := struct {
		invisible   bool
		nameless    bool
		untraceable bool
	}{
		invisible:   false,
		nameless:    true,
		untraceable: true,
	}

	if anon.invisible && anon.nameless && anon.untraceable {
		fmt.Println("You are anonymous.")
	} else {
		fmt.Println("We can find you.")
	}
}
