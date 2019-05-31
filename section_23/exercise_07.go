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
	c := make(chan int)

	for i := 1; i <= 10; i++ {
		go func(k int) {
			for j := 1; j <= 10; j++ {
				c <- k + (j * 10)
				if k * j == 100 {
					close(c)
				}
			}
		}(i)
	}

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("end of program.")
}

