package main

func main() {
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cdards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
func newCard() string {
	return "Five of Diamonds"
}
