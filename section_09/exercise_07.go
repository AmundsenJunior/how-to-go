/*
  Building on the previous hands-on exercise, create a program that uses “else if” and “else”.
  solution: https://play.golang.org/p/IDnrJpE7vT
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	if math.Abs(-4) == 5 {
		fmt.Println("absolute value of -4 is 5")
	} else if 7 == 5 {
		fmt.Println("7 equals 5")
	} else {
		fmt.Println("Do better math")
	}
}
