package main

import "fmt"

/**
 * Main performer.
 */
func perform() {
	cards := newDeck()

	// fmt.Println(cards.toString())

	// hand, remainingDeck := deal(cards, 3)

	// fmt.Println("Hand is =>", hand)
	// fmt.Println("Remaining deck is =>", remainingDeck)
	// remainingDeck.print()
	// cards.print()

	// cards := newDeckFromFile("my_cards")

	// cards.print()

	fmt.Println("------------")
	cards.shuffle()
	cards.print()
}

/**
 * Main function - entry point.
 */
func main() {
	perform()
}
