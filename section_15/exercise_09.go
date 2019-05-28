/*
  A “callback” is when we pass a func into a func as an argument. For this exercise,
  pass a func into a func as an argument
  code: https://play.golang.org/p/0yGYPKh1y7
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

	sleeper := func (f func(string), d time.Duration) func(string) {
		return func(s string) {
			time.Sleep(d * time.Millisecond)
			f(s)
		}
	}

	halfsecond := sleeper(timeLog, 500)
	second := sleeper(timeLog, 1000)

	halfsecond("Starting")
	second("")
	halfsecond("Ending")
}

