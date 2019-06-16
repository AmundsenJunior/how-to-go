/*
  We are going to learn about testing next.
  For this exercise, take a moment and see how much you can figure out about testing
  by reading the testing documentation (http://godoc.org/testing)
  & also Caleb Doxsey’s article on testing (http://www.golang-book.com/books/intro/12).
  See if you can get a basic example of testing working.
*/

package exercise_05

import (
	"fmt"
	"log"
	"math"
)

type SqrtError struct {
	lat  string
	long string
	err  error
}

func (se SqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := Sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, SqrtError{
			lat: "45N",
			long: "110W",
			err: fmt.Errorf("Input %v was less than zero", f),
		}
	}
	return math.Sqrt(f), nil
}
