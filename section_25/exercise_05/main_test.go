

package exercise_05_test

import (
	"github.com/amundsenjunior/how-to-go/section_25/exercise_05"
	"math"
	"testing"
)

var v1, v2, v3 float64 = 15, 22.5, -313.7

func TestSqrt(t *testing.T) {
	val1, err := exercise_05.Sqrt(v1)
	if err != nil {
		t.Errorf("Unexpected test error for square root of %v. Error: %v\n", v1, err)
	}
	exp1 := math.Sqrt(v1)
	if val1 != exp1 {
		t.Errorf("Calculation error for square root of %v. Expected %v, received %v.\n", v1, exp1, val1)
	}

	val2, err := exercise_05.Sqrt(v2)
	if err != nil {
		t.Errorf("Unexpected test error for square root of %v. Error: %v\n", v2, err)
	}
	exp2 := math.Sqrt(v2)
	if val2 != exp2 {
		t.Errorf("Calculation error for square root of %v. Expected %v, received %v.\n", v2, exp2, val2)
	}

	val3, err := exercise_05.Sqrt(v3)
	if err == nil {
		t.Errorf("Sqrt should throw error of type %T on negative argument %v.\n", exercise_05.SqrtError{}, v3)
	}
	var exp3 float64 = 0
	if val3 != exp3 {
		t.Errorf("Calculation error for square root of %v. Expected %v, received %v.\n", v1, exp1, val1)
	}

}

