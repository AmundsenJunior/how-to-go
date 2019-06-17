/*
  Start with this code.
  Get the code ready to BET on (add benchmarks, examples, tests).
  Write a table test.
  Add documentation to the code.
  Run the following in this order:
  tests
  benchmarks
  coverage
  coverage shown in web browser
  examples shown in documentation in a web browser
  helpful to know:
  https://play.golang.org/p/4GUqs1HMpp
  https://play.golang.org/p/P9unTIFeOq
*/

package main

import (
	"fmt"
	"github.com/amundsenjunior/how-to-go/section_29/exercise_03/mymath"
)

func main() {
	xxi := gen()
	for _, v := range xxi {
		fmt.Println(mymath.CenteredAvg(v))
	}
}

func gen() [][]int {
	a := []int{1, 4, 6, 8, 100}
	b := []int{0, 8, 10, 1000}
	c := []int{9000, 4, 10, 8, 6, 12}
	d := []int{123, 744, 140, 200}
	e := [][]int{a, b, c, d}
	return e
}