package main

// Main Caller
func main() {
	// Create a slice
	cards := deck{newCard(), newCard()}
	// Append a new value and return a new slice for assignment.
	cards = append(cards, newCard())

	cards.print()

	newDeck()
}

// Returns a new card name
func newCard() string {
	return "Five of Diamonds"
}
