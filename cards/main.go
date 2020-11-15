package main

import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()

	fmt.Println("\n--- My deck of cards ---")
	cards.print()

	hand, remainingCards := deal(cards, 5)
	fmt.Println("\n--- Cards in my hand ---")
	hand.print()
	fmt.Println("\n--- My remaining cards ---")
	remainingCards.print()

	hand.saveToFile("my_hand")
	remainingCards.saveToFile("my_remainingCards")
}
