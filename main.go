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

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

// Main Caller for using structs
func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Faro",
		contactInfo: contactInfo{
			email:   "jim@example.com",
			zipcode: 12345,
		},
	}

	jim.updateName("Bob")
	jim.print()
}

// Shows pass by reference for persistent value assignment.
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
