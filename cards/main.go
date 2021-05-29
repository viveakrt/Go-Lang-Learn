package main

// import "fmt"

func main() {

	card := openFromFile("CHI")
	// cards := card.toString()
	card.shuffle()
	card.print()
	// fmt.Println(cards)

	// hand, remainCards := deal(card, 5)
	// hand.print()
	// remainCards.print()

}
