/*
  The best way to learn is to teach. For this hands-on exercise,
  choose one of the above exercises, or use the recursion example of factorial
  download, install, and get it running
  https://obsproject.com/
  record a video of YOU teaching the topic
  upload the video to youtube
  share the video on twitter and tag me in it ( https://twitter.com/Todd_McLeod ) so that I can see it!
*/

package main

import "fmt"

func main() {
	f1 := factorial(4)
	fmt.Println("factorial of 4:", f1)
	f2 := factorial(29)
	fmt.Println("factorial of 29:", f2)
}

func factorial(n int) uint64 {
	var factorial uint64 = 1

	for ; n > 0; n-- {
		factorial *= uint64(n)
	}

	return factorial
}
