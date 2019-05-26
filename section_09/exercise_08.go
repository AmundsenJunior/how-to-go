/*
  Create a program that uses a switch statement with no switch expression specified.
  solution: https://play.golang.org/p/YpPgkWGqKY
 */

package main

import "fmt"

func main()  {
	switch {
	case 3 == 2:
		fmt.Println("3 equals 2")
	case !false:
		fmt.Println("not false is true")
	}
}
