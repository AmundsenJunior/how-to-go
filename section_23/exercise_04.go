/*
  Starting with this code(https://play.golang.org/p/MvL6uamrJP),
  pull the values off the channel using a select statement
  solution: https://play.golang.org/p/FulKBY5JNj
*/

package main

import (
	"fmt"
)

func main() {
	c := make(chan int) // action channel
	q := make(chan int) // quit channel

	go func() {
		sender(c, q)
	}()

	receiver(c, q)

	fmt.Println("about to exit")
}

func sender(c, q chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	q <- 1
	close(c)
}

func receiver(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println("value from c channel", v)
		case <-q:
			return
		}
	}
}