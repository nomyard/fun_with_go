package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()

}

//deal returns two values
//first return is assigned to hand
//second return is assigned to remainingCards

// because we are assigning hand and remaining cards to the deal function
// their type is now deck

//since their type is deck we are able to then use the print fucntion
//we built earlier since that function was only able to receive deck type values
