/*
  Using the code from the previous example, delete a record from your map. Now print the map out using the “range” loop
  solution: https://play.golang.org/p/TYl5EbjoeC
*/

package main

import "fmt"

func main() {
	m := map[string][]string{
		"Bond James":      {"Shaken, not stirred", "Martinis", "Women"},
		"Moneypenny Miss": {"James Bond", "Literature", "Computer Science"},
		"No Doctor":       {"Being evil", "Ice cream", "Sunsets"},
	}

	m["Fleming Ian"] = []string{"Writing", "Tea", "Expendable spies"}

	printMap(m)
}

func printMap(m map[string][]string) {
	for k, v := range m {
		fmt.Println(k)
		for idx, val := range v {
			fmt.Printf("\t%v, %v\n", idx, val)
		}
	}
}
