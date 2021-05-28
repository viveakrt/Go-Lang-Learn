package main

import (
	"fmt"
	"strings"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardUnit := deck{"Diamond", "Shade", "Club", "Heart"}
	cardElement := deck{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "King", "Queen", "Joker"}
	for _, cardUnit := range cardUnit {
		for _, cardElement := range cardElement {
			cards = append(cards, cardElement+" Of "+cardUnit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
