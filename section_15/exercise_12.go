/*
  use callback function to extend upon the evenSum code to return the oddSum value
*/

package main

import (
	"fmt"
)

func main() {
	j := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	t := sum(j...)
	fmt.Println(t)
	u := evenSum(sum, j...)
	fmt.Println(u)
	v := oddSum(sum, j...)
	fmt.Println(v)
}

func sum(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}

func evenSum(f func(x ...int) int, y ...int) int {
	var xi []int
	for _, v := range y {
		if v%2 == 0 {
			xi = append(xi, v)
		}
	}
	total := f(xi...)
	return total
}

func oddSum(f func(x ...int) int, y ...int) int {
	var xi []int
	for _, v := range y {
		if v%2 != 0 {
			xi = append(xi, v)
		}
	}
	total := f(xi...)
	return total
}
