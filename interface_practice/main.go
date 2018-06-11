package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//can omit the value call (eb) if not using it
func (englishBot) getGreeting() string {
	return "Hi There!"
}

//can do the same here- but not going to for future reference sake
// func (sb spanishBot) getGreeting() string {
// 	return "Hola!"
// }
