package word

import "strings"

// UseCount returns a map of each unique word in the input string as keys,
// with the number of instances of each word as the value.
// Note that "word" in this function is both case-sensitive
// and includes any special characters adjacent to the alphanumeric word.
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns the number of words in the input string, as separated by a space " " character.
func Count(s string) int {
	sSlice := strings.Split(s, " ")
	return len(sSlice)
}