/*
  Fix the race condition you created in exercise #4 by using package atomic
  code: https://github.com/GoesToEleven/go-programming
*/

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	var increment int64
	const routines = 100

	for i := 0; i < routines; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&increment, 1)
			wg.Done()
		}()

		fmt.Println("increment", atomic.LoadInt64(&increment))
	}

	wg.Wait()
	fmt.Println("increment total:", increment)
}
