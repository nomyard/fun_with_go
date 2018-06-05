package main

import (
	"fmt"
	"io/ioutil"
	"strings"
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

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// d deck

//d is the value, deck is the type

//deck was initialized at the very top of the file as a new slice
// the type of deck is a slice which only receives strings

//the deal function receives the value d (with the type of deck) as an argument
//it also receives the value handSize (with the type of int) as an argument

//we are then returning multiple values. Both return values will be the type of deck
