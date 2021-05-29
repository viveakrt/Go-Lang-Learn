package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func openFromFile(fileName string) deck {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s := strings.Split(string(data), ",")
	return deck(s)
}

func (d deck) shuffle() {
	num := rand.NewSource(time.Now().UnixNano())
	r := rand.New(num)
	for index := range d {
		newPosition := r.Intn(len(d) - 1)

		d[index], d[newPosition] = d[newPosition], d[index]

	}
}
