/*
  Create a func which returns a func
  assign the returned func to a variable
  call the returned func
  code: https://play.golang.org/p/c2HwqVE1Rd
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	timeLog := func(s string) {
		fmt.Printf("%s. Current time of %v\n", s, time.Now())
	}

	sleeper := func (d time.Duration) func() {
		return func() {
			time.Sleep(d * time.Millisecond)
		}
	}

	halfsecond := sleeper(500)
	second := sleeper(1000)

	timeLog("Starting")
	halfsecond()
	timeLog("")
	second()
	timeLog("")
	halfsecond()
	timeLog("Ending")
}


