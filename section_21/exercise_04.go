/*
  Fix the race condition you created in the previous exercise by using a mutex
  it makes sense to remove runtime.Gosched()
  code: https://github.com/GoesToEleven/go-programming
*/

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	var increment int
	const routines = 100

	for i := 0; i < routines; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			increment++
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("increment total:", increment)
}
