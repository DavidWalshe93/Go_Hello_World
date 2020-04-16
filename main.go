package main

// Main Caller
func main() {
	// Create a slice
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()
}
