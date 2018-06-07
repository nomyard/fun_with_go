package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

// This function is turning the deck slice into a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}

//This function is writing to the hard drive
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		//Option #1 - log the error and return a call to newDeck()
		//Option #2 - log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

//go through this function line by line to understand this better!****
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

// d deck

//d is the value, deck is the type

//deck was initialized at the very top of the file as a new slice
// the type of deck is a slice which only receives strings

//the deal function receives the value d (with the type of deck) as an argument
//it also receives the value handSize (with the type of int) as an argument

//we are then returning multiple values. Both return values will be the type of deck
