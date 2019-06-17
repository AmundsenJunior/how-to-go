package main_test

import (
	"github.com/amundsenjunior/how-to-go/section_46/factorial_pipeline"
	"testing"
)

type factorialTest struct {
	number int
	factorial uint64
}

var factorialTests = []factorialTest{
	{3, 6},
	{7, 5040},
	{20, 2432902008176640000},
	{25, 0},
}

func TestFactorial(t *testing.T) {
	for _, test := range factorialTests {
		actual := main.Factorial(test.number)
		if actual != test.factorial {
			t.Errorf("Expected %v, got %v.\n", test.factorial, actual)
		}
	}
}

func BenchmarkFactorial(b *testing.B) {
	for i := 1; i < b.N; i++ {
		main.Factorial(20)
	}
}