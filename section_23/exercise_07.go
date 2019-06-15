/*
  write a program that
  launches 10 goroutines
  each goroutine adds 10 numbers to a channel
  pull the numbers off the channel and print them
  solutions:
  https://play.golang.org/p/R-zqsKS03P
  https://play.golang.org/p/quWnlwzs2z
*/

package main

import (
	"fmt"
)

func main() {
	x := 10
	y := 10
	c := make(chan int)

	for i := 0; i < x; i++ {
		go func(k int) {
			for j := 1; j <= y; j++ {
				x := (k * 10) + j
				c <- x
			}
		}(i)
	}

	for i := 1; i < (x*y); i++ {
		v := <-c
		fmt.Println(v)
	}

	fmt.Println("end of program.")
}

