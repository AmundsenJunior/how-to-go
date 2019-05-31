/*
  Starting with this code (https://play.golang.org/p/sfyu4Is3FG),
  pull the values off the channel using a for range loop
  solution: https://play.golang.org/p/D3N4Tq54SN
*/

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		send(c)
	}()

	receive(c)

	fmt.Println("about to exit")
}

func send(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
