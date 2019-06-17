/*
  From this code (https://github.com/GoesToEleven/GolangTraining/blob/master/22_go-routines/12_channels_pipeline/03_challenge-description/main.go),
  create a pipeline pattern solution for calculating the factorial values of 100 numbers concurrently.
 */

package main

import (
	"fmt"
)

//var factorials = make(map[int]int)

func main() {
	values := gen(20)

	factorials := calc(values)

	for n := range factorials {
		fmt.Println(n)
	}
}

func gen(numValues int) chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= numValues; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func calc(in chan int) chan uint64 {
	out := make(chan uint64)
	go func() {
		for n := range in {
			out <- Factorial(n)
		}
		close(out)
	}()
	return out
}

// Factorial returns the factorial value
// for any number up to 20, or 0 for
// numbers greater than 20
func Factorial(n int) uint64 {
	if n > 20 {
		return 0
	}
	var total uint64 = 1
	for i := n; i > 0; i-- {
		total *= uint64(i)
	}
	return total
}
