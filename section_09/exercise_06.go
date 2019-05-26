/*
  Create a program that shows the “if statement” in action.
  solution: https://play.golang.org/p/DpZ_FLfn5s
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	if math.Abs(-5) == 5 {
		fmt.Println("True")
	}
}
