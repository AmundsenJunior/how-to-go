/*
  get this code (https://play.golang.org/p/j-EA6003P0)) working:
  a buffered channel
  solution: https://play.golang.org/p/Y0Hx6IZc3U
*/

package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
