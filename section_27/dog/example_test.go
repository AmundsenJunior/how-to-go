package dog_test

import (
	"fmt"
	"github.com/amundsenjunior/how-to-go/section_27/dog"
)

func ExampleHumanToDog() {
	years := dog.HumanToDog(10)
	fmt.Println(years)

	// Output: 70
}

func ExampleDogToHuman() {
	years := dog.DogToHuman(35)
	fmt.Println(years)

	// Output: 5
}
