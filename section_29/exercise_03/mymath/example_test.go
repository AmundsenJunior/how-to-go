package mymath_test

import (
	"fmt"
	"github.com/amundsenjunior/how-to-go/section_29/exercise_03/mymath"
)

func ExampleCenteredAvg() {
	average := mymath.CenteredAvg([]int{3, 4, 5, 6, 7})
	fmt.Println(average)

	// Output: 5
}