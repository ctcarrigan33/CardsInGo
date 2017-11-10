package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	// calling from deck.go
	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
