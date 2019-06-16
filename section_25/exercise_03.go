/*
  Create a struct “customErr” which implements the builtin.error interface.
  Create a func “foo” that has a value of type error as a parameter.
  Create a value of type “customErr” and pass it into “foo”.
  If you need a hint, here is one (https://play.golang.org/p/L5QWV8-p11).
  solution:
  https://play.golang.org/p/ixeowY2fd2
  assertion
  https://play.golang.org/p/pbl2kCYsM0
*/

package main

import (
	"fmt"
	"runtime"
)

type customErr struct {
	numGoroutines int
	name string
}

func (c customErr) Error() string {
	return fmt.Sprintf("There were %v goroutines running when object %v was created.\n", c.numGoroutines, c.name)
}

func main() {
	c1 := customErr{
		numGoroutines: runtime.NumGoroutine(),
		name: "c1",
	}

	foo(c1)
}

func foo(e error) {
	fmt.Println(e)
}
