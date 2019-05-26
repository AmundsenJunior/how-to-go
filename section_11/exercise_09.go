/*
  Using the code from the previous example, add a record to your map. Now print the map out using the “range” loop
  solution: https://play.golang.org/p/_CkxAhRrDJ
 */

package main

import "fmt"

func main()  {
	m := map[string][]string {
		"Bond James": {"Shaken, not stirred", "Martinis", "Women"},
		"Moneypenny Miss": {"James Bond", "Literature", "Computer Science"},
		"No Doctor": {"Being evil", "Ice cream", "Sunsets"},
	}

	m["Fleming Ian"] = []string{"Writing", "Tea", "Expendable spies"}

	for k, v := range m {
		fmt.Println(k)
		for idx, val := range v {
			fmt.Printf("\t%v, %v\n", idx, val)
		}
	}
}
