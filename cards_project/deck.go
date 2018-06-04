package main

import (
	"fmt"
)

// Create a new type of 'deck'
//which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{} // cards is initialized to the deck slice

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	//cardSuits is a new slice of values
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	//cardValues is a new slice of values

	for _, suit := range cardSuits {
		//for each suit in the slice of cardSuits
		for _, value := range cardValues {
			//for each value in the slice of cardValues
			cards = append(cards, value+" of "+suit)
			//append to the cards variable, the value and suit to concatenate a string
			//of values to display to the stdout (monitor).
		}
	}
	return cards
	//return the cards variable.
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// type deck []string
// aka deck extends string
//able to use deck in place of []string now

//func (d deck) print() {
// 	for i, card := range d {
// 		fmt.Println(i, card)
// 	}
// }
// (d deck) d is the copy of the cards initialization in main.go.
// deck is referring to the extension from earlier

// for i, card := range d

// that line is making it possible to loop through the cards initialization
// by taking the copy we made (through d) and assigning the contents to
// i, card - we are then able to print out the contents of the original
// cards initialization in main.go
