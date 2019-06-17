package dog_test

import (
	"github.com/amundsenjunior/how-to-go/section_29/exercise_01/dog"
	"testing"
)

func TestYears(t *testing.T) {
	t1 := dog.Years(15)
	if t1 != 105 {
		t.Errorf("Error converting human years to dog years. Expected 105, actual %v\n", t1)
	}
}

func TestYearsTwo(t *testing.T) {
	t2 := dog.YearsTwo(15)
	if t2 != 105 {
		t.Errorf("Error converting human years to dog years. Expected 105, actual %v\n", t2)
	}
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dog.Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dog.YearsTwo(10)
	}
}

