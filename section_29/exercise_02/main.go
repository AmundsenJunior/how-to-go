/*
  Start with this code.
  Get the code ready to BET on (add benchmarks, examples, tests).
  However do not write an example for the func that returns a map, and only write a test for it as an extra challenge.
  Add documentation to the code.
  Run the following in this order:
    tests: go test
    benchmarks: go test -bench .
    coverage: go test -cover
    coverage shown in web browser
    examples shown in documentation in a web browser
*/

package main

import (
	"fmt"
	"github.com/amundsenjunior/how-to-go/section_29/exercise_02/quote"
	"github.com/amundsenjunior/how-to-go/section_29/exercise_02/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}