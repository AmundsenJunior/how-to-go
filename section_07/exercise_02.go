/*
  Using the following operators, write expressions and assign their values to variables:
  ==
  <=
  >=
  !=
  <
  >
  Now print each of the variables.
  solution: https://play.golang.org/p/76R-poSzaY
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	a := math.Abs(-3) == 3
	b := 7 <= 4
	c := 7 >= 4
	d := true != false
	e := "a" > "c"
	f := "a" < "c"
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
	fmt.Println("f", f)
}