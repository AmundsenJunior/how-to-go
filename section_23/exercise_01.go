/*
  get this code (https://play.golang.org/p/j-EA6003P0)) working:
  using func literal, aka, anonymous self-executing func
  solution: https://play.golang.org/p/SHr3lpX4so
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
}
