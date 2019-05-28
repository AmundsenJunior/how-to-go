/*
  Build and use an anonymous func
  code: https://play.golang.org/p/DQX3xEIcRe
*/

package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	x := rand.Float64()

	func(fl float64) {
		sq := math.Sqrt(float64(fl))
		fmt.Printf("the square root of %v is %v\n", fl, sq)
	}(x)
}

