// Package dog provides functions for converting between dog and human years.
package dog

// HumanToDog func returns human years converted into dog years
func HumanToDog(years int) int {
	return years * 7
}

// DogToHuman func returns dog years converted into human years
func DogToHuman(years int) int {
	return years / 7
}