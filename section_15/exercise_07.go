/*
  Assign a func to a variable, then call that func
  code: https://play.golang.org/p/_Qu7ZAyFDH
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	timeLog := func() {
		fmt.Printf("Current time of %v\n", time.Now())
	}

	timeLog()
	time.Sleep(1000 * time.Millisecond)
	timeLog()
}

