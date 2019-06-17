package mymath_test

import (
	"github.com/amundsenjunior/how-to-go/section_29/exercise_03/mymath"
	"math/rand"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type test struct {
		values   []int
		average float64
	}

	tests := []test{
		{[]int{2, 3, 4, 5}, 3.5},
		{[]int{6, 7, 8, 9, 10}, 8},
		{[]int{19, 27, 19, 4, 1}, 14},
		{[]int{-1, 0, 1}, 0},
	}

	for _, v := range tests {
		x := mymath.CenteredAvg(v.values)
		if x != v.average {
			t.Errorf("Expected %v, got %v\n.", v.average, x)
		}
	}
}

func BenchmarkCenteredAvg(b *testing.B) {
	values := make([]int, 10)
	for i := range values {
		values[i] = rand.Intn(100)
	}

	for i := 1; i < b.N; i++ {
		mymath.CenteredAvg(values)
	}
}