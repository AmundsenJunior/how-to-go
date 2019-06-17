/*
  Start with this code.
  Get the code ready to BET on the code (add benchmarks, examples, tests).
  Run the following in this order:
    tests:
      go test -v src/github.com/amundsenjunior/how-to-go/section_29/exercise_01/dog/main_test.go
    benchmarks:
      cd src/github.com/amundsenjunior/how-to-go/section_29/exercise_01/dog
      go test -b .
    coverage:
      cd src/github.com/amundsenjunior/how-to-go/section_29/exercise_01/dog
      go test -cover
    coverage shown in web browser:
      cd src/github.com/amundsenjunior/how-to-go/section_29/exercise_01/dog
      go test -coverprofile c.out
      go tool cover -html=c.out
    examples shown in documentation in a web browser:
      godoc -http=:8080
*/

package main

import (
	"fmt"
	"github.com/amundsenjunior/how-to-go/section_29/exercise_01/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}