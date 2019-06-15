/*
  Starting with this code (https://play.golang.org/p/sfyu4Is3FG),
  pull the values off the channel using a for range loop
  solution: https://play.golang.org/p/D3N4Tq54SN

  This solution shows that we can create and return a pointer to a channel
  within a function called from main, and the goroutine kicked off within
  the gen func still shoots off to run async. So that receives() func
  can start printing off the range of the channel it takes as an arg.
*/

package main

import (
	"fmt"
)

func main() {
	c := gen()
	receives(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	fmt.Println("making channel in gen()")
	c := make(chan int)

	fmt.Println("kicking off goroutine")
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
			fmt.Printf("Sending %v\n", i)
		}
		close(c)
	}()

	fmt.Println("returning channel from gen()")
	return c
}

func receives(c <-chan int) {
  	for v := range c {
		fmt.Printf("Receiving %v\n", v)
	}
}
