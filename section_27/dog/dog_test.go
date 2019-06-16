package dog_test

import (
	"github.com/amundsenjunior/how-to-go/section_27/dog"
	"testing"
)

func TestHumanToDog(t *testing.T) {
	t1 := dog.HumanToDog(15)
	if t1 != 105 {
		t.Errorf("Error converting human years to dog years. Expected 105, actual %v\n", t1)
	}
}

func TestDogToHuman(t *testing.T) {
	t2 := dog.DogToHuman(70)
	if t2 != 10 {
		t.Errorf("Error converting dog years to human years. Expected 10, actual %v\n", t2)
	}
}