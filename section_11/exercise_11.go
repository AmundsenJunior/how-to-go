/*
  once a slice made from make() expands due to append, and the underlying array is extended,
  reducing the size of the slice by itself does not reduce the size of the underlying array back down.
*/

package main

import (
	"fmt"
)

func main() {
	x := make([]int, 5, 10)
	fmt.Println(x)
	fmt.Println(len(x), cap(x))

	x = append(x, x...) // to append a slice to a slice, must unfurl the elements of a slice into the append function
	fmt.Println(x)
	fmt.Println(len(x), cap(x))

	x = append(x, x...)
	fmt.Println(x)
	fmt.Println(len(x), cap(x))

	x = append(x[0:5], 1)
	fmt.Println(x)
	fmt.Println(len(x), cap(x)) // cap does not shrink back down to smaller/original array capacity
}
