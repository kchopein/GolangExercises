package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()

	// cards.saveToFile("cards.deck")

	// cardsFromFile, err := newDeckFromFile("cards.deck")

	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// cardsFromFile.print()

}
