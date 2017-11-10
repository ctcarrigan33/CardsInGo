package main

import "fmt"

// Create new type of deck
// which is a slice of strings

type deck []string

// loop through deck of cards, print values
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
