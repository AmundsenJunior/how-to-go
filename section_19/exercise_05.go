/*
  Starting with this code (https://play.golang.org/p/BVRZTdlUZ_), sort the []user by
  age
  last
  Also sort each []string “Sayings” for each user
  print everything in a way that is pleasant
  solution: https://play.golang.org/p/8RKkdtLl6w
*/

package main

import (
	"fmt"
	"sort"
	"strings"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type users []user

type AgeSort []user
type LastSort []user
type SayingsSort []string

func (u users) printUsers() {
	for _, v := range u {
		fmt.Printf("\t%s %s, age %d\n", v.First, v.Last, v.Age)
		for _, v := range v.Sayings {
			fmt.Printf("\t\t%s\n", v)
		}
	}
}

func (a AgeSort) Len() int {
	return len(a)
}

func (a AgeSort) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func (a AgeSort) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (l LastSort) Len() int {
	return len(l)
}

func (l LastSort) Less(i, j int) bool {
	return l[i].Last < l[j].Last
}

func (l LastSort) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (s SayingsSort) Len() int {
	return len(s)
}

func (s SayingsSort) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s SayingsSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := users{u1, u2, u3}
	fmt.Printf("\n%s\n", strings.ToUpper("Unsorted Users"))
	users.printUsers()

	for _, v := range users {
		sort.Sort(SayingsSort(v.Sayings))
	}
	fmt.Printf("\n%s\n", strings.ToUpper("Sorted Sayings"))
	users.printUsers()

	sort.Sort(AgeSort(users))
	fmt.Printf("\n%s\n", strings.ToUpper("Sorted Ages"))
	users.printUsers()

	sort.Sort(LastSort(users))
	fmt.Printf("\n%s\n", strings.ToUpper("Sorted Last Names"))
	users.printUsers()
}

