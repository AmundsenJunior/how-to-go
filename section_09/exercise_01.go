/*
  Print every number from 1 to 10,000
  solution: https://play.golang.org/p/voDiuiDGGw
*/

package main

import "fmt"

func main() {
	for i := 1; i <= 10000; i++ {
		fmt.Printf("%d\t", i)
	}
}