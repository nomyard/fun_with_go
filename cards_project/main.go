package main

func main() {
	cards := newDeck() //cards is initialized to the newDeck function

	cards.print() // use the print receiver function to display to stdout (monitor)
}
