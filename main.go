// David Walshe
// 16/04/2020

package main

import "fmt"

// Main Caller for Playing Cards
//func main() {
//	// Create a slice
//	cards := PlayingCards.NewDeck()
//
//	cards.SaveToFile("myDeck.txt")
//
//	cards.Shuffle()
//	fmt.Println(cards)
//}

type person struct {
	firstName string
	lastName  string
}

// Main Caller for using structs
func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson"}

	fmt.Println(alex)
}
