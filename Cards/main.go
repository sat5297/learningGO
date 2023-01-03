package main

func main() {
	cards := newDeckFromFile("myCards")
	// cards.print()
	cards.shuffle()
	cards.print()
	// hand, remDeck := deal(cards, 5)
	// fmt.Println(cards.saveToFile("myCards"))
	// fmt.Println()
}
