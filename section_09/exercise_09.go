/*
  Create a program that uses a switch statement with
  the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.
  solution: https://play.golang.org/p/h-8FnjpzWB
*/

package main

import "fmt"

func main() {
	favSport := "pingpong"

	switch favSport {
	case "baseball":
		fmt.Println("my favorite sport is baseball")
	case "skiing":
		fmt.Println("my favorite sport is skiing")
	default:
		fmt.Println("everybody's favorite sport is chess")
	}
}
