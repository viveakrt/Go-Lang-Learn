package main

import "fmt"

func main() {

	card := newDeck()
	cards := card.toString()
	fmt.Println(cards)

	// hand, remainCards := deal(card, 5)
	// hand.print()
	// remainCards.print()

}
