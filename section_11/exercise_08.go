/*
  Create a map with a key of TYPE string which is a person’s “last_first” name,
  and a value of TYPE []string which stores their favorite things.
  Store three records in your map. Print out all of the values, along with their index position in the slice.

  `bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
  `moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
  `no_dr`, `Being evil`, `Ice cream`, `Sunsets`

  solution: https://play.golang.org/p/nTzSlRa9_A
 */

package main

import "fmt"

func main()  {
	m := map[string][]string {
		"Bond James": {"Shaken, not stirred", "Martinis", "Women"},
		"Moneypenny Miss": {"James Bond", "Literature", "Computer Science"},
		"No Doctor": {"Being evil", "Ice cream", "Sunsets"},
	}

	for k, v := range m {
		fmt.Println(k)
		for idx, val := range v {
			fmt.Printf("\t%v, %v\n", idx, val)
		}
	}
}
