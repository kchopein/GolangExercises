package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) (deck, error) {
	fileContent, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	fileAsStrings := strings.Split(string(fileContent), ",")

	return deck(fileAsStrings), err
}

func (d deck) shuffle() {
	randSource := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(randSource)

	for i := range d {
		newPosition := randomGenerator.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
