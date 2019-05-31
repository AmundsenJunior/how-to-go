/*
  Get this code running:
  https://play.golang.org/p/_DBRueImEq
  solution: https://play.golang.org/p/mgw750EPp4
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
	fmt.Printf("cs\t%T\n", c)
}
