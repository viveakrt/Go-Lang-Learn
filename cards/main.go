package main

func main() {

	card := newDeck()
	hand, remainCards := deal(card, 5)
	hand.print()
	remainCards.print()

}
