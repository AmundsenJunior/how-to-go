package dog_test

import (
	"fmt"
	"github.com/amundsenjunior/how-to-go/section_29/exercise_01/dog"
)

func ExampleYears() {
	years := dog.Years(10)
	fmt.Println(years)

	// Output: 70
}
