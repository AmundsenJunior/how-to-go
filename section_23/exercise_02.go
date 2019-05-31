/*
  Get this code running:
  https://play.golang.org/p/oB-p3KMiH6
  solution: https://play.golang.org/p/isnJ8hMMKg
*/

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()
	fmt.Println(<-c)

	fmt.Printf("------\n")
	fmt.Printf("c\t%T\n", c)
}
