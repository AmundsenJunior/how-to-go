/*
  in addition to the main goroutine, launch two additional goroutines
  each additional goroutine should print something out
  use waitgroups to make sure each goroutine finishes before your program exists
  code: https://github.com/GoesToEleven/go-programming
*/

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func one() {
	fmt.Println("First function executing...")
	wg.Done()
}

func two() {
	fmt.Println("Second function executing...")
	wg.Done()
}

func main() {
	fmt.Println("Main function starting...")

	wg.Add(1)
	go one()

	wg.Add(1)
	go two()

	wg.Wait()
	fmt.Println("Main function ending...")
}

