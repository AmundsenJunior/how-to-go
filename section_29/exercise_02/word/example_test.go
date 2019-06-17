package word_test

import (
	"fmt"
	"github.com/amundsenjunior/how-to-go/section_29/exercise_02/word"
)

func ExampleCount() {
	s := "one two three four"
	fmt.Println(word.Count(s))

	// Output: 4
}
